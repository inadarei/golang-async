package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/buger/jsonparser"
)

const baseURL = "https://www.googleapis.com/books/v1/volumes?q="

type fnBytes func([]byte)

func main() {
	defer panicHandler()

	type PublishRecord struct {
		author string
		count  int
	}
	var record PublishRecord
	var records = []PublishRecord{}

	const isbn = "1491956224" // "Microservice Architecture"
	authors := bookAuthors(isbn)
	// fmt.Println("Authors: %v", authors)

	var wg sync.WaitGroup
	wg.Add(len(authors))

	for _, author := range authors {
		go func(anAuthor string) {
			defer panicHandler()
			defer wg.Done()

			numBooks, authorName := authorNumBooks(anAuthor)
			record.author = authorName
			record.count = numBooks
			records = append(records, record)
		}(author)
	}

	wg.Wait()

	fmt.Printf("Publishing statistics: %+v \n", records)
}

/** panicHandler handles all the panics. If you need stack trace, uncomment the
  corresponding line */
func panicHandler() {
	if r := recover(); r != nil {
		// fmt.Printf("ERROR: %v\n\n %s", r, debug.Stack())
		fmt.Printf("ERROR: %v \n", r)
		os.Exit(1)
	}
}

/** bookAuthors returns all authors of a book */
func bookAuthors(isbn string) []string {
	bookURL := baseURL + "isbn:" + isbn

	body := queryAPI(bookURL, googleAPIError)

	authors := []string{}

	// items[0].volumeInfo.authors;
	jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		valueString := string(value[:])
		authors = append(authors, valueString)
	}, "items", "[0]", "volumeInfo", "authors")

	if len(authors) == 0 {
		panic(fmt.Sprintf("no authors found for isbn: %v, with URL: %v", isbn, bookURL))
	}
	return authors
}

/** authorNumBooks returns number of books an author has published */
func authorNumBooks(authorName string) (int, string) {
	authorNameSafe := url.QueryEscape(authorName)
	apiURL := baseURL + "inauthor:\"" + authorNameSafe + "\""

	body := queryAPI(apiURL, googleAPIError)

	numBooks, err := jsonparser.GetInt(body, "totalItems")
	if err != nil {
		panic(err.Error())
	}

	return int(numBooks), authorName
}

/** queryAPI an HTTP API with error handling */
func queryAPI(apiURL string, errHandler fnBytes) []byte {
	res, err := http.Get(apiURL)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode > 399 {
		errHandler(body)
	}

	return body
}

/** googleAPIError handles any non-OK HTTP status returns */
func googleAPIError(body []byte) {
	errMsg, err := jsonparser.GetString(body, "error", "message")
	if err != nil {
		panic(err.Error())
	}
	panic(fmt.Sprintf("Google Books API Response: '%s'", errMsg))
}
