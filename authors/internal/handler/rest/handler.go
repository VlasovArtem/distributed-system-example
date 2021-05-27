package rest

import (
	"encoding/json"
	"fmt"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type invalidRequestError struct {
	code int
	message string
}

func newError(code int, message string) error {
	return &invalidRequestError{
		code:    code,
		message: message,
	}
}

func (i *invalidRequestError) Bytes() []byte  {
	return []byte(i.message)
}

func (i *invalidRequestError) Error() string {
	return fmt.Sprintf("Code: %d. Message: %s", i.code, i.message)
}

func New(s *service.Service) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/authors", FindAllAuthors(s)).Methods("GET")
	router.HandleFunc("/api/v1/authors/{id}", FindAuthorById(s)).Methods("GET")
	router.HandleFunc("/api/v1/authors", AddAuthor(s)).Methods("POST")
	router.HandleFunc("/api/v1/authors", UpdateAuthor(s)).Methods("PUT")

	return router
}

func UpdateAuthor(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author, err := readAuthorFromBody(r)

		if err != nil {
			requestError := err.(*invalidRequestError)
			w.WriteHeader(requestError.code)
			w.Write(requestError.Bytes())
			return
		}

		existsById := s.ExistsById(author.ID)

		if !existsById {
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Author with id %s is not found", author.ID)))
			return
		}

		s.Save(author)
	}
}

func AddAuthor(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author, err := readAuthorFromBody(r)

		if err != nil {
			requestError := err.(*invalidRequestError)
			w.WriteHeader(requestError.code)
			w.Write(requestError.Bytes())
			return
		}

		existsById := s.ExistsById(author.ID)

		if existsById {
			w.WriteHeader(400)
			w.Write([]byte(fmt.Sprintf("Author with id %s is exists.", author.ID)))
			return
		}

		s.Save(author)
	}
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

func readAuthorFromBody(r *http.Request) (author service.Author, err error) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return author, newError(400, "Could not read body")
	}

	err = json.Unmarshal(body, &author)

	if err != nil {
		return author, newError(400, "Could not parse body")
	}

	return author, nil
}
