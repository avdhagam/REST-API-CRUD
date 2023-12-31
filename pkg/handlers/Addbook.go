package handlers

import (
	"REST-API-CRUD/pkg/mocks"
	"REST-API-CRUD/pkg/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	//Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	//Append to Book mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)
	//Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
