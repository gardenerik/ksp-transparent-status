package candle

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type LessonTime int

var Lessons = []LessonTime{
	810,
	900,
	950,
	1040,
	1130,
	1220,
	1310,
	1400,
	1450,
	1540,
	1630,
	1720,
	1810,
	1900,
}

type CandleXML struct {
	XMLName xml.Name       `xml:"timetable"`
	Lessons []CandleLesson `xml:"lesson"`
}

type CandleTimeString string

func (s CandleTimeString) LessonTime() LessonTime {
	parts := strings.SplitN(string(s), ":", 2)
	if len(parts) != 2 {
		return 0
	}

	h, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}

	m, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}

	return LessonTime(h * 100 + m)
}

type CandleLesson struct {
	XMLName xml.Name `xml:"lesson"`
	Id      int      `xml:"id,attr"`
	Type    string   `xml:"type"`
	Room    string   `xml:"room"`
	Subject string   `xml:"subject"`
	Day     string   `xml:"day"`
	Start   CandleTimeString   `xml:"start"`
	End     CandleTimeString   `xml:"end"`
	Teacher string   `xml:"teacher"`
	Note    string   `xml:"note"`
}

func (l CandleLesson) Weekday() time.Weekday {
	switch l.Day {
	case "Po":
		return time.Monday
	case "Ut":
		return time.Tuesday
	case "St":
		return time.Wednesday
	case "Št":
		return time.Thursday
	case "Pi":
		return time.Friday
	}

	return time.Sunday
}
