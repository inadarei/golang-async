package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/buger/jsonparser"
)

func main() {
	var authors []string
	const isbn = "1491956224" // "Microservice Architecture"
	authors = bookAuthors(isbn)

	fmt.Printf("%v", authors)

}

/* bookAuthors returns all authors of a book */
func bookAuthors(isbn string) []string {
	baseURL := "https://www.googleapis.com/books/v1/volumes?q=isbn:"
	bookURL := baseURL + isbn

	res, err := http.Get(bookURL)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	authors := []string{}

	// items[0].volumeInfo.authors;
	jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		valueString := string(value[:])
		authors = append(authors, valueString)
	}, "items", "[0]", "volumeInfo", "authors")

	return authors
}

/* authorBooks returns number of books an author has published */
func authorNumBooks() {

}
