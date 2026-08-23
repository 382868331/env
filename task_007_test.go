package env

import "testing"

func TestTask007RequiredMissing(t *testing.T) {
	var v struct {
		Name string `env:"NAME,required"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{}})
	if e == nil {
		t.Fatal("missing required variable returned nil")
	}
}
