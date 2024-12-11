package bookstore

import (
	"errors"
	"fmt"
)

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

type Book struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	Copies          int    `json:"copies"`
	ID              int
	PriceCents      int `json:"net_price_cents"`
	DiscountPercent int
	category        Category
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

func (b *Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(p int) error {
	if p < 0 {
		return errors.New("price cents must be greater than zero")
	}

	(*b).PriceCents = p
	return nil
}

func (b *Book) SetCategory(c Category) error {

	if !validCategory[c] {
		return fmt.Errorf("unknown category %q", c)
	}

	b.category = c
	return nil
}

func (b Book) Category() Category {
	return b.category
}
