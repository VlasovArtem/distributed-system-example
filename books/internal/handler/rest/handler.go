package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/VlasovArtem/distributed-system-example/books/internal/service"
)

func New(s *service.Service) http.Handler {
	// TODO create 2 handlers: api/v1/books and api/v1/books/{id}
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/books", FindAllBooks(s))
	router.HandleFunc("/api/v1/books/{id}", FindBookById(s))

	return router
}

func FindBookById(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := vars["id"]

		book, err := s.GetByID(id)

		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Book with %s is not found", id)))
		} else {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func FindAllBooks(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books := s.List()

		json.NewEncoder(w).Encode(books)
	}
}
