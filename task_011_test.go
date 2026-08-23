package env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTask011ReadsFileContent(t *testing.T) {
	p := filepath.Join(t.TempDir(), "value")
	if e := os.WriteFile(p, []byte("secret"), 0600); e != nil {
		t.Fatal(e)
	}
	v, e := getFromFile(p)
	if e != nil {
		t.Fatal(e)
	}
	if v != "secret" {
		t.Fatalf("value=%q", v)
	}
}
func TestTask011KeepsNewline(t *testing.T) {
	p := filepath.Join(t.TempDir(), "value")
	if e := os.WriteFile(p, []byte("a\nb"), 0600); e != nil {
		t.Fatal(e)
	}
	v, e := getFromFile(p)
	if e != nil {
		t.Fatal(e)
	}
	if v != "a\nb" {
		t.Fatalf("value=%q", v)
	}
}
