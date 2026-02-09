package mytypes_test

import (
	"testing"
	"mytypes"
)

func TestTwice (t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("Twice: %v, want %v, got %v", input, want, got)
	}
}

func TestMyStringLen (t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("Hello")
	want := 5
	got := input.Len()
	if want != got {
		t.Errorf("Length: %v, want %v, got %v", input, want, got)
	}
}