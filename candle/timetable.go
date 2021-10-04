package candle

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
	"zahradnik.xyz/mirror-stats/config"
)

type CachedTimetable struct {
	Lessons     []CandleLesson
	RetrievedAt time.Time
}

var TimetableCache = map[string]CachedTimetable{}

func LoadTimetable(name string) ([]CandleLesson, error) {
	resp, err := http.Get(fmt.Sprintf("https://candle.fmph.uniba.sk/rozvrh/%v.xml", name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %v while downloading RozvrhXML", resp.StatusCode)
	}

	var data CandleXML
	err = xml.NewDecoder(resp.Body).Decode(&data)
	return data.Lessons, err
}

func GetTimetable(name string) CachedTimetable {
	timetable, exists := TimetableCache[name]
	if !exists || time.Now().Sub(timetable.RetrievedAt) >= 6*time.Hour {
		tt, err := LoadTimetable(name)
		if err != nil {
			log.Println(err)
			return CachedTimetable{}
		}
		dat := CachedTimetable{
			Lessons:     tt,
			RetrievedAt: time.Now(),
		}
		TimetableCache[name] = dat
		return dat
	}

	return timetable
}

func GetPeopleHavingLesson(l LessonTime, day time.Weekday) []config.Person {
	out := []config.Person{}

	for _, person := range config.People {
		tt := GetTimetable(person.CandleName)
		for _, lesson := range tt.Lessons {
			if lesson.Weekday() != day {
				continue
			}

			if lesson.Start.LessonTime() <= l && lesson.End.LessonTime() >= l {
				out = append(out, person)
				break
			}
		}
	}

	return out
}
