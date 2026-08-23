package env

import (
	"reflect"
	"testing"
)

func TestTask017SliceOptionsKeepFuncMap(t *testing.T) {
	fm := map[reflect.Type]ParserFunc{reflect.TypeOf(0): func(string) (interface{}, error) { return 1, nil }}
	o := optionsWithSliceEnvPrefix(Options{FuncMap: fm}, 0)
	if len(o.FuncMap) != 1 {
		t.Fatalf("FuncMap length=%d", len(o.FuncMap))
	}
}
func TestTask017SliceOptionsKeepEnvironment(t *testing.T) {
	e := map[string]string{"A": "b"}
	o := optionsWithSliceEnvPrefix(Options{Environment: e}, 1)
	if o.Environment["A"] != "b" {
		t.Fatalf("Environment=%v", o.Environment)
	}
}
