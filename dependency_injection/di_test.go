package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Minh")

	got := buffer.String()
	want := "Hello, Minh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
