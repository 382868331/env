package env

import "testing"

func TestTask012FirstSlicePrefix(t *testing.T) {
	o := optionsWithSliceEnvPrefix(Options{Prefix: "ITEM_"}, 0)
	if o.Prefix != "ITEM_0_" {
		t.Fatalf("Prefix=%q", o.Prefix)
	}
}
func TestTask012LaterSlicePrefix(t *testing.T) {
	o := optionsWithSliceEnvPrefix(Options{Prefix: "ITEM_"}, 3)
	if o.Prefix != "ITEM_3_" {
		t.Fatalf("Prefix=%q", o.Prefix)
	}
}
