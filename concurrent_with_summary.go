package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"github.com/buger/jsonparser"
)

const baseURL = "https://www.googleapis.com/books/v1/volumes?q="

func main() {
	//var authors []string
	const isbn = "1491956224" // "Microservice Architecture"
	authors := bookAuthors(isbn)

	var wg sync.WaitGroup
	wg.Add(len(authors))

	for _, author := range authors {
		go func(author string) {
			defer wg.Done()

			numBooks, authorName := authorNumBooks(author)
			fmt.Printf("%s authored %d books \n", authorName, numBooks)
		}(author)
	}

	wg.Wait()
}

/* bookAuthors returns all authors of a book */
func bookAuthors(isbn string) []string {
	bookURL := baseURL + "isbn:" + isbn

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

/* authorBooks returns books an author has published */
func authorNumBooks(authorName string) (int, string) {
	authorNameSafe := url.QueryEscape(authorName)
	apiURL := baseURL + "inauthor:\"" + authorNameSafe + "\""

	res, err := http.Get(apiURL)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	numBooks, err := jsonparser.GetInt(body, "totalItems")
	if err != nil {
		panic(err.Error())
	}

	return int(numBooks), authorName
}
