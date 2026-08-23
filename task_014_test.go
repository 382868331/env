package env

import "testing"

func TestTask014PreservesTagKey(t *testing.T) {
	k, o := parseKeyForOption("API_TOKEN,required")
	if k != "API_TOKEN" || len(o) != 1 || o[0] != "required" {
		t.Fatalf("key=%q opts=%v", k, o)
	}
}
func TestTask014MixedCaseKey(t *testing.T) {
	k, _ := parseKeyForOption("MyValue")
	if k != "MyValue" {
		t.Fatalf("key=%q", k)
	}
}
