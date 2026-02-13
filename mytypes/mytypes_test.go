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


func TestMyBuilderHello (t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "Hello, Farhan!"
	got := mb.Hello()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestStringBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Farhan")
	want := "Hello, Farhan"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := mb.Contents.Len()
	if wantLen != gotLen {
		t.Errorf("%q wantLen %v, gotLen %v", mb.Contents.String(), wantLen, gotLen)
	}
}

func TestStringUpperCaser(t *testing.T) {
	t.Parallel()
	var su mytypes.MyBuilder
	su.Contents.WriteString("hello, farhan")
	want := "HELLO, FARHAN"
	got := su.ToUpper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}