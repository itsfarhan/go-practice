package mytypes

import "strings"

type MyInt int

type MyString string

type MyBuilder struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) Len() int {
	return len(s)
}

func (m MyBuilder) Hello() string {
	return "Hello, Farhan!"
}
