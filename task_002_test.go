package env

import (
	"testing"
	"time"
)

func TestTask002PositiveDuration(t *testing.T) {
	v, e := parseDuration("1500ms")
	if e != nil {
		t.Fatal(e)
	}
	if v.(time.Duration) != 1500*time.Millisecond {
		t.Fatalf("duration=%v", v)
	}
}
