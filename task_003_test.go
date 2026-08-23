package env

import (
	"testing"
	"time"
)

func TestTask003NamedLocation(t *testing.T) {
	v, e := parseLocation("Asia/Tokyo")
	if e != nil {
		t.Fatal(e)
	}
	loc := v.(time.Location)
	if loc.String() != "Asia/Tokyo" {
		t.Fatalf("location=%s", loc.String())
	}
}
func TestTask003UTCLocation(t *testing.T) {
	v, e := parseLocation("UTC")
	if e != nil {
		t.Fatal(e)
	}
	loc := v.(time.Location)
	if loc.String() != "UTC" {
		t.Fatalf("location=%s", loc.String())
	}
}
