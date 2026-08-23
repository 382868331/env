package env

import "testing"

func TestTask015StandardTag(t *testing.T) {
	var v struct {
		Name string `env:"NAME"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"NAME": "alice"}})
	if e != nil {
		t.Fatal(e)
	}
	if v.Name != "alice" {
		t.Fatalf("Name=%q", v.Name)
	}
}
