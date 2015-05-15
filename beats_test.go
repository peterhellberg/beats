package beats_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/peterhellberg/beats"
)

func TestBeatsString(t *testing.T) {
	in, out := 14, "@014"

	if got := beats.String(in); got != out {
		t.Errorf("String(%v) = %v, want %v", in, got, out)
	}
}

func TestBeatsFromTime(t *testing.T) {
	in, out := time.Date(2014, 9, 2, 9, 23, 20, 0, time.UTC), 432

	if got := beats.FromTime(in); got != out {
		t.Errorf("FromTime(%v) = %v, want %v", in, got, out)
	}
}

func ExampleFromTime() {
	fmt.Println(beats.FromTime(time.Date(2015, 5, 14, 12, 35, 43, 0, time.UTC)))
	// Output: 566
}

func TestBeatsBMT(t *testing.T) {
	in, out := time.Date(2012, 4, 11, 16, 45, 51, 0, time.UTC),
		time.Date(2012, 4, 11, 17, 45, 51, 0, time.UTC)

	if got := beats.BMT(in); got != out {
		t.Errorf("BMT(%v) = %v, want %v", in, got, out)
	}
}
