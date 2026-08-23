package env

import "testing"

func TestTask005ExistingSeparator(t *testing.T) {
	if got := toEnvName("API_Key"); got != "API_KEY" {
		t.Fatalf("got=%q", got)
	}
}
