package beats

import (
	"testing"
	"time"
)

func TestBeatsString(t *testing.T) {
	in, out := 14, "@014"

	if got := String(in); got != out {
		t.Errorf("String(%v) = %v, want %v", in, got, out)
	}
}

func TestBeatsFromTime(t *testing.T) {
	in, out := time.Date(2014, 9, 2, 9, 23, 20, 0, time.UTC), 432

	if got := FromTime(in); got != out {
		t.Errorf("FromTime(%v) = %v, want %v", in, got, out)
	}
}

func TestBeatsBMT(t *testing.T) {
	in, out := time.Date(2012, 4, 11, 16, 45, 51, 0, time.UTC),
		time.Date(2012, 4, 11, 17, 45, 51, 0, time.UTC)

	if got := BMT(in); got != out {
		t.Errorf("BMT(%v) = %v, want %v", in, got, out)
	}
}
