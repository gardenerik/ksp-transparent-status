package clock

import (
	"github.com/nsf/termbox-go"
	"time"
)

type SquarePosition struct {
	x, y int
}

var SQUARES = []SquarePosition{
	{0, 0}, {2, 0}, {4, 0}, {0, 1}, {4, 1}, {0, 2}, {2, 2}, {4, 2}, {0, 3}, {4, 3}, {0, 4}, {2, 4}, {4, 4},
}

var NUMBERS = [][]int{
	{0, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12},
	{2, 4, 7, 9, 12},
	{0, 1, 2, 4, 5, 6, 7, 8, 10, 11, 12},
	{0, 1, 2, 4, 5, 6, 7, 9, 10, 11, 12},
	{0, 2, 3, 4, 5, 6, 7, 9, 12},
	{0, 1, 2, 3, 5, 6, 7, 9, 10, 11, 12},
	{0, 1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12},
	{0, 1, 2, 4, 7, 9, 12},
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	{0, 1, 2, 3, 4, 5, 6, 7, 9, 12},
}

func PrintTime() {
	t := time.Now()
	bg := termbox.RGBToAttribute(255, 255, 255)

	drawNumber(5, 2, t.Hour()/10, bg)
	drawNumber(13, 2, t.Hour()%10, bg)

	drawSquare(13+8, 3, bg)
	drawSquare(13+8, 5, bg)

	drawNumber(13+12, 2, t.Minute()/10, bg)
	drawNumber(13+12+8, 2, t.Minute()%10, bg)
}

func drawNumber(x, y, num int, bg termbox.Attribute) {
	for _, sq := range NUMBERS[num] {
		square := SQUARES[sq]
		drawSquare(x+square.x, y+square.y, bg)
	}
}

func drawSquare(x, y int, bg termbox.Attribute) {
	termbox.SetBg(x, y, bg)
	termbox.SetBg(x+1, y, bg)
}
