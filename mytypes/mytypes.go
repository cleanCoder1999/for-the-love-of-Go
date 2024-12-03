package mytypes

import "strings"

type MyString string

func (s MyString) Len() int {
	return len(s)
}

type MyBuilder struct {
	Contents strings.Builder
}

func (b MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

type StringUpperCaser struct {
	Contents strings.Builder
}

func (b StringUpperCaser) ToUpper() string {
	return strings.ToUpper(b.Contents.String())
}

func Double(x *int) {
	*x *= 2
}

type MyInt int

func (mi *MyInt) Double() {
	*mi *= 2
}
