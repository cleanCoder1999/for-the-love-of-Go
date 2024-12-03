package mytypes_test

import (
	"mytypes"
	"strings"
	"testing"
)

func TestMyStringLen(t *testing.T) {
	t.Parallel()

	s := "Hello World"
	myStr := mytypes.MyString(s)

	l := myStr.Len()
	if l != len(s) {
		t.Errorf("want %d, got %d", len(s), l)
	}
}

func TestStringsBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder

	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	wantLen := 15
	gotLen := sb.Len()

	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", sb.String(),
			wantLen, gotLen)
	}
}

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestMyBuilder(t *testing.T) {
	t.Parallel()

	var mb mytypes.MyBuilder

	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")

	want := "Hello, Gophers!"
	got := mb.Contents.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	wantLen := 15
	gotLen := mb.Contents.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d",
			mb.Contents.String(), wantLen, gotLen)
	}
}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()

	var tu mytypes.StringUpperCaser

	tu.Contents.WriteString("Hello, ")
	tu.Contents.WriteString("Gophers!")

	want := "HELLO, GOPHERS!"
	got := tu.ToUpper()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()

	x := 12
	want := 24

	mytypes.Double(&x)

	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}

func TestDoubleP(t *testing.T) {
	t.Parallel()

	x := mytypes.MyInt(12)
	want := mytypes.MyInt(24)

	p := &x
	p.Double()

	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}
