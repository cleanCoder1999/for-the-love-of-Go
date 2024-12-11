package creditcard

import "errors"

type card struct {
	number string
}

func (c *card) Number() string {
	return c.number
}

func New(number string) (card, error) {
	if number == "" {
		return card{}, errors.New("number must not be empty")
	}

	return card{number: number}, nil
}
