package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New(booksDest, authorsDest string) http.Handler {
	// TODO create handler api/v1/dashboard
	router := mux.NewRouter().StrictSlash(true)
	return router
}
