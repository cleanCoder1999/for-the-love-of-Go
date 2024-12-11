package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestCreditCard(t *testing.T) {
	t.Parallel()

	want := "AT11 2000 1234 5678 9091"

	card, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}

	got := card.Number()
	if got == "" {
		t.Errorf("got %q; want non-empty number", got)
	}

	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

func TestNewInvalidReturnsError(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")
	if err == nil {
		t.Fatal("want error for invalid card number, got nil")
	}
}
