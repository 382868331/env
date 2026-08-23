package env

import "testing"

func TestTask012FirstSlicePrefix(t *testing.T) {
	o := optionsWithSliceEnvPrefix(Options{Prefix: "ITEM_"}, 0)
	if o.Prefix != "ITEM_0_" {
		t.Fatalf("Prefix=%q", o.Prefix)
	}
}
