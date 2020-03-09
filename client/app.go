package main

import (
	"log"
	"net/rpc"
	
	"github.com/ridwanakf/nadc-intro-to-rpc/client/internal"
)

func main() {
	var reply internal.Book
	var db []internal.Book

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("[Connection error] ", err)
	}

	book1 := internal.Book{
		ID:    1,
		Title: "The Alchemist",
	}
	book2 := internal.Book{
		ID:    2,
		Title: "We were The Lucky Ones",
	}
	book3 := internal.Book{
		ID:    1,
		Title: "Looking for Alaska",
	}

	err = client.Call("API.AddNewBook", book1, &reply)
	if err != nil {
		log.Printf("error when calling API.AddNewBook: %+v", err)
		return
	}

	err = client.Call("API.AddNewBook", book2, &reply)
	if err != nil {
		log.Printf("error when calling API.AddNewBook: %+v", err)
		return
	}

	err = client.Call("API.AddNewBook", book3, &reply)
	if err != nil {
		log.Printf("error when calling API.AddNewBook: %+v", err)
		return
	}

	err = client.Call("API.GetDB", "", &db)
	if err != nil {
		log.Printf("error when calling API.GetDB: %+v", err)
		return
	}

	log.Printf("Database client: %+v\n", db)
}
