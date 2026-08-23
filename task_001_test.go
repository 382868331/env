package env

import (
	"net/url"
	"testing"
)

func TestTask001URLComponents(t *testing.T) {
	v, e := parseURL("https://example.com/a?q=1")
	if e != nil {
		t.Fatal(e)
	}
	u := v.(url.URL)
	if u.Scheme != "https" || u.Host != "example.com" || u.Path != "/a" {
		t.Fatalf("url=%#v", u)
	}
}
