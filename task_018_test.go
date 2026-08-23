package env

import "testing"

func TestTask018DefaultFillsZero(t *testing.T) {
	v := struct {
		Port int `env:"PORT" envDefault:"8080"`
	}{}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{}, SetDefaultsForZeroValuesOnly: true})
	if e != nil {
		t.Fatal(e)
	}
	if v.Port != 8080 {
		t.Fatalf("Port=%d", v.Port)
	}
}
func TestTask018DefaultPreservesNonZero(t *testing.T) {
	v := struct {
		Port int `env:"PORT" envDefault:"8080"`
	}{Port: 9000}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{}, SetDefaultsForZeroValuesOnly: true})
	if e != nil {
		t.Fatal(e)
	}
	if v.Port != 9000 {
		t.Fatalf("Port=%d", v.Port)
	}
}
