package env

import "testing"

func TestTask006EmptyUsesDefault(t *testing.T) {
	v, ok, d := getOr("PORT", "8080", true, map[string]string{"PORT": ""})
	if v != "8080" || !ok || !d {
		t.Fatalf("v=%q ok=%v default=%v", v, ok, d)
	}
}
