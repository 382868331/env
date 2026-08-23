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
