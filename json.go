package main

import (
	"encoding/json"
	"fmt"
)

type (
	//Book struct
	Book struct {
		ISBN   string `json:"isbn"`
		Title  string `json:"title"`
		Author Author `json:"author"`
	}

	//Author struct
	Author struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
)

//ToJSON converts a struct to JSON
func ToJSON() {
	author := Author{
		FirstName: "Hank",
		LastName:  "Moody",
	}

	book := Book{
		ISBN:   "0123456789",
		Title:  "God Hates Us All",
		Author: author,
	}

	bytes, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

//FullName of the author
func (a *Author) FullName() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

//FromJSON hydrates a struct from JSON
func FromJSON() {
	js := `
	{
		"isbn": "0123456789",
		"title": "God Hates Us All",
		"author": {
			"firstName": "Hank",
			"lastName": "Moody"
		}
	}`

	var book Book
	json.Unmarshal([]byte(js), &book)
	fmt.Printf("%s by %s, ISBN %s", book.Title, book.Author.FullName(), book.ISBN)
}
