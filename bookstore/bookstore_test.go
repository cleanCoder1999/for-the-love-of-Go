package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"sort"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}

	want := 1

	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	got := catalog.GetAllBooks()

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}

	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()

	want := 7500

	book := bookstore.Book{
		ID:              1,
		Title:           "For the Love of Go",
		PriceCents:      10000,
		DiscountPercent: 25,
	}

	got := book.NetPriceCents()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	want := 500

	book := bookstore.Book{
		ID:              1,
		Title:           "For the Love of Go",
		PriceCents:      10000,
		DiscountPercent: 25,
	}

	err := book.SetPriceCents(want)
	if err != nil {
		t.Fatal("want no error, got ", err)
	}

	got := book.PriceCents

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{}
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}

	for _, cat := range cats {
		err := b.SetCategory(cat)
		if err != nil {
			t.Fatal(err)
		}
		got := b.Category()
		if cat != got {
			t.Errorf("want category %q, got %q", cat, got)
		}
	}
}

// ########### INVALID INPUT CASES
func TestGetBookInvalid(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	_, err := catalog.GetBook(3)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		ID:              1,
		Title:           "For the Love of Go",
		PriceCents:      10000,
		DiscountPercent: 25,
	}

	err := book.SetPriceCents(-500)
	if err == nil {
		t.Fatal("want error, got nil")
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory(999)
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}
