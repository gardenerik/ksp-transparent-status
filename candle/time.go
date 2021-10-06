package candle

import (
	"fmt"
	"time"
)

func CandleTime() LessonTime {
	t := time.Now()
	return LessonTime(t.Hour()*100 + t.Minute())
}

func (t LessonTime) String() string {
	return fmt.Sprintf("%d:%02d", t/100, t%100)
}

func GetCurrentLesson() (LessonTime, bool) {
	candleTime := CandleTime()
	previous := LessonTime(0)
	for _, hour := range Lessons {
		if hour > candleTime {
			return previous, previous != 0
		}
		previous = hour
	}
	return previous, previous != 0
}

func GetNextLesson() (LessonTime, bool) {
	candleTime := CandleTime()
	for _, hour := range Lessons {
		if hour >= candleTime {
			return hour, true
		}
	}
	return 0, false
}
