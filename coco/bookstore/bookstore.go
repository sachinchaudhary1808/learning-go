// Package bookstore pkg
package bookstore

import "errors"

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

// Buy function
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}
