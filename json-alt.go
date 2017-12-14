// This sample demonstrates decoding JSON strings in concurrent_with_summary.go
// without using an external package or precise mapping with a predefined struct
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	authors := getAuthors(bookJSON)
	fmt.Printf("Authors: %v \n", authors)

	numBooks := getNumBooks(authorJSON)
	fmt.Printf("Ronnie has published %v books\n", numBooks)

}

// items[0].volumeInfo.authors;
func getAuthors(jsonChunk []byte) []string {
	var jsonData map[string]interface{}
	err := json.Unmarshal(jsonChunk, &jsonData)
	if err != nil {
		panic(err)
	}

	var items = jsonData["items"].([]interface{})
	var authors []interface{}
	for _, item := range items {
		volumeInfo := item.(map[string]interface{})["volumeInfo"]
		authorsArr := volumeInfo.(map[string]interface{})["authors"]
		authors = authorsArr.([]interface{})
	}

	// Convert elements to be strings
	var strAuthors []string
	for _, author := range authors {
		strAuthors = append(strAuthors, author.(string))
	}

	return strAuthors
}

func getNumBooks(jsonChunk []byte) int {
	var jsonData map[string]interface{}
	err := json.Unmarshal(jsonChunk, &jsonData)
	if err != nil {
		panic(err)
	}

	var maybeNum = jsonData["totalItems"] // this is interface
	num := int(maybeNum.(float64))        // convert to float with assertion and then to int

	return num
}

//------- Sample JSON response payloads

// sample response
var bookJSON = []byte(`{
	"kind": "books#volumes",
	"totalItems": 1,
	"items": [
	 {
		"kind": "books#volume",
		"id": "gf2wDAAAQBAJ",
		"etag": "mEk52ofi1Yw",
		"selfLink": "https://www.googleapis.com/books/v1/volumes/gf2wDAAAQBAJ",
		"volumeInfo": {
		 "title": "Microservice Architecture",
		 "subtitle": "Aligning Principles, Practices, and Culture",
		 "authors": [
			"Irakli Nadareishvili",
			"Ronnie Mitra",
			"Matt McLarty",
			"Mike Amundsen"
		 ],
		 "publisher": "\"O'Reilly Media, Inc.\"",
		 "publishedDate": "2016-07-18",
		 "description": "Microservices can have a positive impact on your enterprise—just ask Amazon and Netflix—but you can fall into many traps if you don’t approach them in the right way. This practical guide covers the entire microservices landscape, including the principles, technologies, and methodologies of this unique, modular style of system building. You’ll learn about the experiences of organizations around the globe that have successfully adopted microservices. In three parts, this book explains how these services work and what it means to build an application the Microservices Way. You’ll explore a design-based approach to microservice architecture with guidance for implementing various elements. And you’ll get a set of recipes and practices for meeting practical, organizational, and cultural challenges to microservice adoption. Learn how microservices can help you drive business objectives Examine the principles, practices, and culture that define microservice architectures Explore a model for creating complex systems and a design process for building a microservice architecture Learn the fundamental design concepts for individual microservices Delve into the operational elements of a microservices architecture, including containers and service discovery Discover how to handle the challenges of introducing microservice architecture in your organization",
		 "industryIdentifiers": [
			{
			 "type": "ISBN_13",
			 "identifier": "9781491956229"
			},
			{
			 "type": "ISBN_10",
			 "identifier": "1491956224"
			}
		 ],
		 "readingModes": {
			"text": true,
			"image": true
		 },
		 "pageCount": 146,
		 "printType": "BOOK",
		 "categories": [
			"Computers"
		 ],
		 "maturityRating": "NOT_MATURE",
		 "allowAnonLogging": true,
		 "contentVersion": "1.3.3.0.preview.3",
		 "panelizationSummary": {
			"containsEpubBubbles": false,
			"containsImageBubbles": false
		 },
		 "imageLinks": {
			"smallThumbnail": "http://books.google.com/books/content?id=gf2wDAAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
			"thumbnail": "http://books.google.com/books/content?id=gf2wDAAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
		 },
		 "language": "en",
		 "previewLink": "http://books.google.com/books?id=gf2wDAAAQBAJ&printsec=frontcover&dq=isbn:1491956224&hl=&cd=1&source=gbs_api",
		 "infoLink": "https://play.google.com/store/books/details?id=gf2wDAAAQBAJ&source=gbs_api",
		 "canonicalVolumeLink": "https://market.android.com/details?id=book-gf2wDAAAQBAJ"
		},
		"saleInfo": {
		 "country": "US",
		 "saleability": "FOR_SALE",
		 "isEbook": true,
		 "listPrice": {
			"amount": 33.99,
			"currencyCode": "USD"
		 },
		 "retailPrice": {
			"amount": 18.35,
			"currencyCode": "USD"
		 },
		 "buyLink": "https://play.google.com/store/books/details?id=gf2wDAAAQBAJ&rdid=book-gf2wDAAAQBAJ&rdot=1&source=gbs_api",
		 "offers": [
			{
			 "finskyOfferType": 1,
			 "listPrice": {
				"amountInMicros": 3.399E7,
				"currencyCode": "USD"
			 },
			 "retailPrice": {
				"amountInMicros": 1.835E7,
				"currencyCode": "USD"
			 },
			 "giftable": true
			}
		 ]
		},
		"accessInfo": {
		 "country": "US",
		 "viewability": "PARTIAL",
		 "embeddable": true,
		 "publicDomain": false,
		 "textToSpeechPermission": "ALLOWED",
		 "epub": {
			"isAvailable": true
		 },
		 "pdf": {
			"isAvailable": true
		 },
		 "webReaderLink": "http://play.google.com/books/reader?id=gf2wDAAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
		 "accessViewStatus": "SAMPLE",
		 "quoteSharingAllowed": false
		},
		"searchInfo": {
		 "textSnippet": "You’ll learn about the experiences of organizations around the globe that have successfully adopted microservices. In three parts, this book explains how these services work and what it means to build an application the Microservices Way."
		}
	 }
	]
 }`)

var authorJSON = []byte(`{
	"kind": "books#volumes",
	"totalItems": 4,
	"items": [
	 {
		"kind": "books#volume",
		"id": "Ev2wDAAAQBAJ",
		"etag": "QcnhFXJLyik",
		"selfLink": "https://www.googleapis.com/books/v1/volumes/Ev2wDAAAQBAJ",
		"volumeInfo": {
		 "title": "Microservice Architecture",
		 "subtitle": "Aligning Principles, Practices, and Culture",
		 "authors": [
			"Irakli Nadareishvili",
			"Ronnie Mitra",
			"Matt McLarty",
			"Mike Amundsen"
		 ],
		 "publisher": "\"O'Reilly Media, Inc.\"",
		 "publishedDate": "2016-07-18",
		 "description": "Microservices can have a positive impact on your enterprise—just ask Amazon and Netflix—but you can fall into many traps if you don’t approach them in the right way. This practical guide covers the entire microservices landscape, including the principles, technologies, and methodologies of this unique, modular style of system building. You’ll learn about the experiences of organizations around the globe that have successfully adopted microservices. In three parts, this book explains how these services work and what it means to build an application the Microservices Way. You’ll explore a design-based approach to microservice architecture with guidance for implementing various elements. And you’ll get a set of recipes and practices for meeting practical, organizational, and cultural challenges to microservice adoption. Learn how microservices can help you drive business objectives Examine the principles, practices, and culture that define microservice architectures Explore a model for creating complex systems and a design process for building a microservice architecture Learn the fundamental design concepts for individual microservices Delve into the operational elements of a microservices architecture, including containers and service discovery Discover how to handle the challenges of introducing microservice architecture in your organization",
		 "industryIdentifiers": [
			{
			 "type": "ISBN_13",
			 "identifier": "9781491956342"
			},
			{
			 "type": "ISBN_10",
			 "identifier": "1491956348"
			}
		 ],
		 "readingModes": {
			"text": false,
			"image": true
		 },
		 "pageCount": 146,
		 "printType": "BOOK",
		 "categories": [
			"Computers"
		 ],
		 "maturityRating": "NOT_MATURE",
		 "allowAnonLogging": false,
		 "contentVersion": "0.3.0.0.preview.1",
		 "panelizationSummary": {
			"containsEpubBubbles": false,
			"containsImageBubbles": false
		 },
		 "imageLinks": {
			"smallThumbnail": "http://books.google.com/books/content?id=Ev2wDAAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
			"thumbnail": "http://books.google.com/books/content?id=Ev2wDAAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
		 },
		 "language": "en",
		 "previewLink": "http://books.google.com/books?id=Ev2wDAAAQBAJ&printsec=frontcover&dq=inauthor:Ronnie+Mitra&hl=&cd=1&source=gbs_api",
		 "infoLink": "http://books.google.com/books?id=Ev2wDAAAQBAJ&dq=inauthor:Ronnie+Mitra&hl=&source=gbs_api",
		 "canonicalVolumeLink": "https://books.google.com/books/about/Microservice_Architecture.html?hl=&id=Ev2wDAAAQBAJ"
		},
		"saleInfo": {
		 "country": "US",
		 "saleability": "NOT_FOR_SALE",
		 "isEbook": false
		},
		"accessInfo": {
		 "country": "US",
		 "viewability": "PARTIAL",
		 "embeddable": true,
		 "publicDomain": false,
		 "textToSpeechPermission": "ALLOWED",
		 "epub": {
			"isAvailable": false
		 },
		 "pdf": {
			"isAvailable": false
		 },
		 "webReaderLink": "http://play.google.com/books/reader?id=Ev2wDAAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
		 "accessViewStatus": "SAMPLE",
		 "quoteSharingAllowed": false
		},
		"searchInfo": {
		 "textSnippet": "You’ll learn about the experiences of organizations around the globe that have successfully adopted microservices. In three parts, this book explains how these services work and what it means to build an application the Microservices Way."
		}
	 },
	 {
		"kind": "books#volume",
		"id": "6ZrEAgAAQBAJ",
		"etag": "8lwuIzgF9kw",
		"selfLink": "https://www.googleapis.com/books/v1/volumes/6ZrEAgAAQBAJ",
		"volumeInfo": {
		 "title": "DataPower SOA Appliance Administration, Deployment, and Best Practices",
		 "authors": [
			"Gerry Kaplan",
			"Jan Bechtold",
			"Daniel Dickerson",
			"Richard Kinard",
			"Ronnie Mitra",
			"Helio L. P. Mota",
			"David Shute",
			"John Walczyk",
			"IBM Redbooks"
		 ],
		 "publisher": "IBM Redbooks",
		 "publishedDate": "2011-06-06",
		 "description": "This IBM® Redbooks® publication focuses on operational and managerial aspects for DataPower® appliance deployments. DataPower appliances provide functionality that crosses both functional and organizational boundaries, which introduces unique management and operational challenges. For example, a DataPower appliance can provide network functionality, such as load balancing, and at the same time, provide enterprise service bus (ESB) capabilities, such as transformation and intelligent content-based routing. This IBM Redbooks publication provides guidance at both a general and technical level for individuals who are responsible for planning, installation, development, and deployment. It is not intended to be a \"how-to\" guide, but rather to help educate you about the various options and methodologies that apply to DataPower appliances. In addition, many chapters provide a list of suggestions.",
		 "industryIdentifiers": [
			{
			 "type": "ISBN_13",
			 "identifier": "9780738435701"
			},
			{
			 "type": "ISBN_10",
			 "identifier": "0738435708"
			}
		 ],
		 "readingModes": {
			"text": true,
			"image": true
		 },
		 "pageCount": 300,
		 "printType": "BOOK",
		 "categories": [
			"Computers"
		 ],
		 "maturityRating": "NOT_MATURE",
		 "allowAnonLogging": true,
		 "contentVersion": "0.3.4.0.preview.3",
		 "panelizationSummary": {
			"containsEpubBubbles": false,
			"containsImageBubbles": false
		 },
		 "imageLinks": {
			"smallThumbnail": "http://books.google.com/books/content?id=6ZrEAgAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
			"thumbnail": "http://books.google.com/books/content?id=6ZrEAgAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
		 },
		 "language": "en",
		 "previewLink": "http://books.google.com/books?id=6ZrEAgAAQBAJ&pg=PP1&dq=inauthor:Ronnie+Mitra&hl=&cd=2&source=gbs_api",
		 "infoLink": "https://play.google.com/store/books/details?id=6ZrEAgAAQBAJ&source=gbs_api",
		 "canonicalVolumeLink": "https://market.android.com/details?id=book-6ZrEAgAAQBAJ"
		},
		"saleInfo": {
		 "country": "US",
		 "saleability": "FOR_SALE",
		 "isEbook": true,
		 "listPrice": {
			"amount": 0.0,
			"currencyCode": "USD"
		 },
		 "retailPrice": {
			"amount": 0.0,
			"currencyCode": "USD"
		 },
		 "buyLink": "https://play.google.com/store/books/details?id=6ZrEAgAAQBAJ&rdid=book-6ZrEAgAAQBAJ&rdot=1&source=gbs_api"
		},
		"accessInfo": {
		 "country": "US",
		 "viewability": "ALL_PAGES",
		 "embeddable": true,
		 "publicDomain": false,
		 "textToSpeechPermission": "ALLOWED",
		 "epub": {
			"isAvailable": true
		 },
		 "pdf": {
			"isAvailable": true
		 },
		 "webReaderLink": "http://play.google.com/books/reader?id=6ZrEAgAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
		 "accessViewStatus": "SAMPLE",
		 "quoteSharingAllowed": false
		},
		"searchInfo": {
		 "textSnippet": "Gerry Kaplan, Jan Bechtold, Daniel Dickerson, Richard Kinard, Ronnie \u003cb\u003eMitra\u003c/b\u003e, \u003cbr\u003e\nHelio L. P. Mota, David Shute, John Walczyk, IBM Redbooks. DataPower SOA \u003cbr\u003e\nAppliance Administration, Deployment, and Best Practices Demonstrates user \u003cbr\u003e\nadministration and role-based management Explains network configuration, \u003cbr\u003e\nmonitoring, and logging Describes appliance and configuration management \u003cbr\u003e\nFront cover Gerry Kaplan Jan Bechtold Daniel Dickerson Richard Kinard Ronnie \u003cbr\u003e\n\u003cb\u003eMitra\u003c/b\u003e Helio L. P.&nbsp;..."
		}
	 },
	 {
		"kind": "books#volume",
		"id": "couaAQAACAAJ",
		"etag": "+8Wbl0F7/Ok",
		"selfLink": "https://www.googleapis.com/books/v1/volumes/couaAQAACAAJ",
		"volumeInfo": {
		 "title": "Microservice Architecture",
		 "authors": [
			"Irakli Nadareishvili. Ronnie Mitra. Matt McLarty. Mike Amundsen"
		 ],
		 "publishedDate": "2016",
		 "industryIdentifiers": [
			{
			 "type": "ISBN_10",
			 "identifier": "1491956321"
			},
			{
			 "type": "ISBN_13",
			 "identifier": "9781491956328"
			}
		 ],
		 "readingModes": {
			"text": false,
			"image": false
		 },
		 "printType": "BOOK",
		 "maturityRating": "NOT_MATURE",
		 "allowAnonLogging": false,
		 "contentVersion": "preview-1.0.0",
		 "panelizationSummary": {
			"containsEpubBubbles": false,
			"containsImageBubbles": false
		 },
		 "language": "en",
		 "previewLink": "http://books.google.com/books?id=couaAQAACAAJ&dq=inauthor:Ronnie+Mitra&hl=&cd=3&source=gbs_api",
		 "infoLink": "http://books.google.com/books?id=couaAQAACAAJ&dq=inauthor:Ronnie+Mitra&hl=&source=gbs_api",
		 "canonicalVolumeLink": "https://books.google.com/books/about/Microservice_Architecture.html?hl=&id=couaAQAACAAJ"
		},
		"saleInfo": {
		 "country": "US",
		 "saleability": "NOT_FOR_SALE",
		 "isEbook": false
		},
		"accessInfo": {
		 "country": "US",
		 "viewability": "NO_PAGES",
		 "embeddable": false,
		 "publicDomain": false,
		 "textToSpeechPermission": "ALLOWED",
		 "epub": {
			"isAvailable": false
		 },
		 "pdf": {
			"isAvailable": false
		 },
		 "webReaderLink": "http://play.google.com/books/reader?id=couaAQAACAAJ&hl=&printsec=frontcover&source=gbs_api",
		 "accessViewStatus": "NONE",
		 "quoteSharingAllowed": false
		}
	 },
	 {
		"kind": "books#volume",
		"id": "nQ2AAAAAMAAJ",
		"etag": "Bti+QozU98A",
		"selfLink": "https://www.googleapis.com/books/v1/volumes/nQ2AAAAAMAAJ",
		"volumeInfo": {
		 "title": "Pariwisata",
		 "subtitle": "antara obsesi dan realita",
		 "authors": [
			"Ronnie Sugiantoro Viko"
		 ],
		 "publishedDate": "2000",
		 "description": "The future of tourist trade in Yogyakarta; collection of articles previously published.",
		 "industryIdentifiers": [
			{
			 "type": "ISBN_10",
			 "identifier": "9799246423"
			},
			{
			 "type": "ISBN_13",
			 "identifier": "9789799246424"
			}
		 ],
		 "readingModes": {
			"text": false,
			"image": false
		 },
		 "pageCount": 155,
		 "printType": "BOOK",
		 "categories": [
			"Tourism"
		 ],
		 "maturityRating": "NOT_MATURE",
		 "allowAnonLogging": false,
		 "contentVersion": "preview-1.0.0",
		 "imageLinks": {
			"smallThumbnail": "http://books.google.com/books/content?id=nQ2AAAAAMAAJ&printsec=frontcover&img=1&zoom=5&source=gbs_api",
			"thumbnail": "http://books.google.com/books/content?id=nQ2AAAAAMAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api"
		 },
		 "language": "id",
		 "previewLink": "http://books.google.com/books?id=nQ2AAAAAMAAJ&q=inauthor:Ronnie+Mitra&dq=inauthor:Ronnie+Mitra&hl=&cd=4&source=gbs_api",
		 "infoLink": "http://books.google.com/books?id=nQ2AAAAAMAAJ&dq=inauthor:Ronnie+Mitra&hl=&source=gbs_api",
		 "canonicalVolumeLink": "https://books.google.com/books/about/Pariwisata.html?hl=&id=nQ2AAAAAMAAJ"
		},
		"saleInfo": {
		 "country": "US",
		 "saleability": "NOT_FOR_SALE",
		 "isEbook": false
		},
		"accessInfo": {
		 "country": "US",
		 "viewability": "NO_PAGES",
		 "embeddable": false,
		 "publicDomain": false,
		 "textToSpeechPermission": "ALLOWED",
		 "epub": {
			"isAvailable": false
		 },
		 "pdf": {
			"isAvailable": false
		 },
		 "webReaderLink": "http://play.google.com/books/reader?id=nQ2AAAAAMAAJ&hl=&printsec=frontcover&source=gbs_api",
		 "accessViewStatus": "NONE",
		 "quoteSharingAllowed": false
		},
		"searchInfo": {
		 "textSnippet": "Kompetitor pun bisa memahami dan akan sama-sama fight bila memang ingin \u003cbr\u003e\ntetap survive. Namun lain masalah tatkala pola manajemen bisnis ini diterapkan \u003cbr\u003e\nuntuk pengelolaan objek wisata. Ternyata, benturan bukan pada pelaku bisnis \u003cbr\u003e\nsejenis yang jadi kompetitornya, namun justru pada supplier-nya, dalam hal ini \u003cbr\u003e\n\u003cb\u003emitra\u003c/b\u003e pendukung yang sekaligus memberikan kontribusi mendatangkan tamu. \u003cbr\u003e\nContoh kasus adalah ketika PT Taman Wisata Candi Borobudur, Candi \u003cbr\u003e\nPrambanan, dan&nbsp;..."
		}
	 }
	]
 }`)
