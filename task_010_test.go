package env

import "testing"

func TestTask010DefaultCallbackFlag(t *testing.T) {
	flag := false
	var v struct {
		Name string `env:"NAME" envDefault:"guest"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{}, OnSet: func(_ string, _ interface{}, d bool) { flag = d }})
	if e != nil {
		t.Fatal(e)
	}
	if !flag {
		t.Fatal("default callback flag=false")
	}
}
func TestTask010ExplicitCallbackFlag(t *testing.T) {
	flag := true
	var v struct {
		Name string `env:"NAME" envDefault:"guest"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"NAME": "alice"}, OnSet: func(_ string, _ interface{}, d bool) { flag = d }})
	if e != nil {
		t.Fatal(e)
	}
	if flag {
		t.Fatal("explicit callback flag=true")
	}
}
