package candle

import (
	"encoding/xml"
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

type CandleLesson struct {
	XMLName xml.Name `xml:"lesson"`
	Id      int      `xml:"id,attr"`
	Type    string   `xml:"type"`
	Room    string   `xml:"room"`
	Subject string   `xml:"subject"`
	Day     string   `xml:"day"`
	Start   string   `xml:"start"`
	End     string   `xml:"end"`
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
