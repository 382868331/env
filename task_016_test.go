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
