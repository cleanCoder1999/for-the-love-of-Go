package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	Copies          int    `json:"copies"`
	ID              int
	PriceCents      int `json:"net_price_cents"`
	DiscountPercent int
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {

	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	books := make([]Book, len(c))

	var j int
	for i := range c {
		books[j] = c[i]
		j++
	}

	return books
}

func (c Catalog) GetBook(id int) (Book, error) {

	if book, ok := c[id]; ok {
		return book, nil
	}

	return Book{}, fmt.Errorf("no book with id=%d found", id)
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}
