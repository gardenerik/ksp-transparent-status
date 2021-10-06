package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"os/signal"
	"syscall"
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

	BindSIGHUP()

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

func BindSIGHUP () {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func(){
		for range c {
			err := config.ReadConfig()
			if err != nil {
				panic(err)
			}

			termbox.Clear(termbox.ColorDefault, termbox.RGBToAttribute(255, 0, 0))
			termbox.Flush()
		}
	}()
}