package module_one

import "testing"

func TestOne(t *testing.T) {
	want := "Hello, i'm the module one!"
	if got := One(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
