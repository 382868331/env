package env

import "testing"

func TestTask004ValueContainsEquals(t *testing.T) {
	m := toMap([]string{"TOKEN=a=b=c"})
	if m["TOKEN"] != "a=b=c" {
		t.Fatalf("TOKEN=%q", m["TOKEN"])
	}
}
func TestTask004EmptyValue(t *testing.T) {
	m := toMap([]string{"EMPTY="})
	v, ok := m["EMPTY"]
	if !ok || v != "" {
		t.Fatalf("EMPTY=%q present=%v", v, ok)
	}
}
