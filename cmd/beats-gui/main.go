package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
	"github.com/peterhellberg/beats"
)

func main() {
	err := ui.Main(func() {
		w := ui.NewWindow("Beats GUI", 205, 24, false)
		w.SetMargined(true)

		l := ui.NewLabel("Time in beats:")
		w.SetChild(l)

		w.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		go func() {
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()

			for now := range ticker.C {
				ui.QueueMain(func() {
					l.SetText(timeInBeatsString(now))
				})
			}
		}()

		w.Show()
	})
	if err != nil {
		panic(err)
	}
}

func timeInBeatsString(t time.Time) string {
	f := "Time in beats: %v (%02d:%02d:%02d)"

	return fmt.Sprintf(f, beats.NowString(), t.Hour(), t.Minute(), t.Second())
}
