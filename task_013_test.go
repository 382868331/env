package env

import (
	"reflect"
	"testing"
)

func TestTask013ComposesNestedPrefix(t *testing.T) {
	type C struct {
		Inner struct{} `envPrefix:"DB_"`
	}
	f, _ := reflect.TypeOf(C{}).FieldByName("Inner")
	o := optionsWithEnvPrefix(f, Options{Prefix: "APP_", PrefixTagName: "envPrefix"})
	if o.Prefix != "APP_DB_" {
		t.Fatalf("Prefix=%q", o.Prefix)
	}
}
