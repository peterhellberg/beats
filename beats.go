/*

A package for all your Swatch Internet Time needs.

http://en.wikipedia.org/wiki/Swatch_Internet_Time

Installation

    go get -u github.com/peterhellberg/beats

Usage

		package main

		import (
			"fmt"

			"github.com/peterhellberg/beats"
		)

		func main() {
			fmt.Println(beats.NowString())
		}
*/
package beats

import (
	"fmt"
	"time"
)

// Now returns the current time in beats
func Now() int {
	return FromTime(time.Now())
}

// NowString returns the current time in beats as a string
func NowString() string {
	return String(Now())
}

// String returns the given time in beats as a string
func String(b int) string {
	return fmt.Sprintf("@%03d", b)
}

// FromTime returns the beats at the given time
func FromTime(t time.Time) int {
	bmt := BMT(t)

	return ((bmt.Hour()*60+bmt.Minute())*60 + bmt.Second()) * 10 / 864
}

// BMT returns the given time in Biel Mean Time
func BMT(t time.Time) time.Time {
	return t.UTC().Add(3600 * time.Second)
}
