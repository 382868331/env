package env

import "testing"

func TestTask020CallbackRuns(t *testing.T) {
	calls := 0
	var v struct {
		Name string `env:"NAME"`
	}
	e := ParseWithOptions(&v, Options{Environment: map[string]string{"NAME": "alice"}, OnSet: func(string, interface{}, bool) { calls++ }})
	if e != nil {
		t.Fatal(e)
	}
	if calls != 1 {
		t.Fatalf("calls=%d", calls)
	}
}
