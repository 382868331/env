package env

import "testing"

func TestTask009ExpandsReference(t *testing.T) {
	var v struct {
		URL string `env:"URL,expand"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"HOST": "example.com", "URL": "https://${HOST}/api"}})
	if e != nil {
		t.Fatal(e)
	}
	if v.URL != "https://example.com/api" {
		t.Fatalf("URL=%q", v.URL)
	}
}
func TestTask009PlainValue(t *testing.T) {
	var v struct {
		URL string `env:"URL,expand"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"URL": "local"}})
	if e != nil || v.URL != "local" {
		t.Fatalf("URL=%q err=%v", v.URL, e)
	}
}
