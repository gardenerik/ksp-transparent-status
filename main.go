package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"time"
	"zahradnik.xyz/mirror-stats/candle"
	"zahradnik.xyz/mirror-stats/clock"
	"zahradnik.xyz/mirror-stats/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	err = termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	termbox.SetOutputMode(termbox.OutputRGB)
	termbox.SetInputMode(termbox.InputEsc)

	go func() {
		for {
			if e := termbox.PollEvent(); e.Type == termbox.EventKey && e.Key == termbox.KeyEsc {
				termbox.Close()
				os.Exit(0)
			}
		}
	}()

	for {
		termbox.Clear(termbox.ColorDefault, termbox.RGBToAttribute(0, 0, 0))
		clock.PrintTime()

		candle.RenderTimetables()

		termbox.Flush()
		time.Sleep(time.Second)
	}
}
