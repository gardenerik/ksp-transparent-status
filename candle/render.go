package candle

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"time"
)

func RenderTimetables() {
	if CandleTime() >= 1950 {
		return
	}

	x := 45
	y := 2

	current, exists := GetCurrentLesson()
	if exists {
		end := header(x, y, termbox.RGBToAttribute(0, 0, 0)|termbox.AttrBold, termbox.RGBToAttribute(76, 209, 55), " Na hodine: ")
		header(end+1, y, termbox.RGBToAttribute(200, 200, 200), termbox.RGBToAttribute(0, 0, 0), current.String())

		y = printlist(x, y+1, GetPeopleHavingLesson(current, time.Now().Weekday()))
		y += 2
	}

	next, exists := GetNextLesson()
	if exists {
		end := header(x, y, termbox.RGBToAttribute(0, 0, 0)|termbox.AttrBold, termbox.RGBToAttribute(251, 197, 49), " Ďaľšia hodina: ")
		header(end+1, y, termbox.RGBToAttribute(200, 200, 200), termbox.RGBToAttribute(0, 0, 0), next.String())

		y = printlist(x, y+1, GetPeopleHavingLesson(next, time.Now().Weekday()))
	}
}

func header(x, y int, fg, bg termbox.Attribute, msg string) int {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
	return x
}

func printlist(x, y int, items []CandlePerson) int {
	w, _ := termbox.Size()
	startX := x
	for i, person := range items {
		item := person.DisplayName
		if i != len(items)-1 {
			item += ", "
		}
		totalW := 0
		for _, c := range item {
			totalW += runewidth.RuneWidth(c)
		}

		if x+totalW > w-5 {
			x = startX
			y++
		}

		attr := person.Color()
		if person.IsOnline {
			attr = attr | termbox.AttrUnderline
		}

		for _, c := range item {
			termbox.SetCell(x, y, c, attr, termbox.RGBToAttribute(0, 0, 0))
			x += runewidth.RuneWidth(c)
		}
	}

	return y
}
