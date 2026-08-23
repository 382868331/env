package env

import (
	"os"
	"testing"
)

func TestTask019UnsetAfterRead(t *testing.T) {
	const k = "TASK019_SECRET"
	t.Setenv(k, "secret")
	var v struct {
		Secret string `env:"TASK019_SECRET,unset"`
	}
	if e := Parse(&v); e != nil {
		t.Fatal(e)
	}
	if _, ok := os.LookupEnv(k); ok {
		t.Fatal("variable still present")
	}
}
