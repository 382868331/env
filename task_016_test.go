package env

import "testing"

func TestTask016ExpansionUsesPriorField(t *testing.T) {
	var v struct {
		Host string `env:"HOST" envDefault:"service.local"`
		URL  string `env:"URL,expand"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"URL": "https://${HOST}/v1"}})
	if e != nil {
		t.Fatal(e)
	}
	if v.URL != "https://service.local/v1" {
		t.Fatalf("URL=%q", v.URL)
	}
}
func TestTask016NestedExpansion(t *testing.T) {
	var v struct {
		Base string `env:"BASE" envDefault:"api"`
		URL  string `env:"URL,expand"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"URL": "${BASE}.example"}})
	if e != nil || v.URL != "api.example" {
		t.Fatalf("URL=%q err=%v", v.URL, e)
	}
}
