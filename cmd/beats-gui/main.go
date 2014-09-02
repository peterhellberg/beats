package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
	"github.com/peterhellberg/beats"
)

var (
	l = ui.NewStandaloneLabel("Time in beats:")
	w = ui.NewWindow("Beats GUI", 205, 24, l)
)

func main() {
	go ui.Do(initGUI)

	go func() {
		c := time.Tick(500 * time.Millisecond)

		for now := range c {
			l.SetText(timeInBeatsString(now))
		}
	}()

	err := ui.Go()
	if err != nil {
		panic(err)
	}
}

func initGUI() {
	w.OnClosing(func() bool {
		ui.Stop()
		return true
	})

	w.Show()
}

func timeInBeatsString(t time.Time) string {
	f := "Time in beats: %v (%02d:%02d:%02d)"

	return fmt.Sprintf(f, beats.NowString(), t.Hour(), t.Minute(), t.Second())
}
