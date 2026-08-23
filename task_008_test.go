package env

import "testing"

func TestTask008RejectsEmpty(t *testing.T) {
	var v struct {
		Name string `env:"NAME,notEmpty"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"NAME": ""}})
	if e == nil {
		t.Fatal("empty value accepted")
	}
}
func TestTask008AcceptsNonEmpty(t *testing.T) {
	var v struct {
		Name string `env:"NAME,notEmpty"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"NAME": "x"}})
	if e != nil || v.Name != "x" {
		t.Fatalf("Name=%q err=%v", v.Name, e)
	}
}
