package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.lohika.com/dmiroshnichenko/distributed-comm-stubs/books/internal/service"
)

func New(s *service.Service) http.Handler {
	// TODO create 2 handlers: api/v1/books and api/v1/books/{id}
	router := mux.NewRouter().StrictSlash(true)
	return router
}
