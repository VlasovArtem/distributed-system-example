package rest

import (
	"encoding/json"
	"fmt"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)



func New(s *service.Service) http.Handler {
	// TODO create 2 handlers: api/v1/authors and api/v1/authors/{id}
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/authors", FindAllAuthors(s))
	router.HandleFunc("/api/v1/authors/{id}", FindAuthorById(s))

	return router
}

func FindAuthorById(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := vars["id"]

		author, err := s.GetByID(id)

		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Author with %s is not found", id)))
		} else {
			json.NewEncoder(w).Encode(author)
		}
	}
}

func FindAllAuthors(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors := s.List()

		json.NewEncoder(w).Encode(authors)
	}
}
