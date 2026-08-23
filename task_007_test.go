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
func TestTask007RequiredPresent(t *testing.T) {
	var v struct {
		Name string `env:"NAME,required"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"NAME": "ok"}})
	if e != nil || v.Name != "ok" {
		t.Fatalf("Name=%q err=%v", v.Name, e)
	}
}
