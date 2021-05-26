package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"

	"github.com/VlasovArtem/distributed-system-example/books/internal/service"
)

type invalidRequestError struct {
	code int
	message []byte
}

func newError(code int, message string) error {
	return &invalidRequestError{
		code:    code,
		message: []byte(message),
	}
}

func newErrorByte(code int, message []byte) error {
	return &invalidRequestError{
		code:    code,
		message: message,
	}
}

func (i *invalidRequestError) Bytes() []byte  {
	return i.message
}

func (i *invalidRequestError) Error() string {
	return fmt.Sprintf("Code: %d. Message: %s", i.code, i.message)
}

func New(s *service.Service) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/books", FindAllBooks(s))
	router.HandleFunc("/api/v1/books/{id}", FindBookById(s))
	router.HandleFunc("/api/v1/books", AddBook(s)).Methods("POST")

	return router
}

func AddBook(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book, err := readBody(r)

		if err != nil {
			requestError := err.(*invalidRequestError)
			w.WriteHeader(requestError.code)
			w.Write(requestError.Bytes())
			return
		}

		existsById := s.ExistsById(book.ID)

		if existsById {
			w.WriteHeader(400)
			w.Write([]byte(fmt.Sprintf("Book with id %s is exists.", book.ID)))
			return
		}

		err = updateNumberOfBooks(s.Cfg.AuthorsHTTP.URL, book.AuthorID)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
		} else {
			s.Save(book)
		}
	}
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

func readBody(r *http.Request) (book service.Book, err error) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return book, newError(400, "Could not read body")
	}

	err = json.Unmarshal(body, &book)

	if err != nil {
		return book, newError(400, "Could not parse body")
	}

	return book, nil
}

func updateNumberOfBooks(authorHttpURL string, authorId int) error {
	authorMap, err := findAuthor(authorHttpURL, authorId)

	if err != nil {
		return err
	}

	authorMap["NumberOfBooks"] = authorMap["NumberOfBooks"].(uint) + 1

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(authorMap)

	client := http.DefaultClient

	request, err := http.NewRequest("PUT", authorHttpURL, b)

	if err != nil {
		return newError(400, "Could not create request to update author")
	}

	request.Header.Set("Content-Type", "application/json")

	if response, err := client.Do(request); err != nil || response.StatusCode != 200 {
		return newError(400, "Request to the author service is not performed.")
	}

	return nil
}

func findAuthor(authorHttpURL string, authorId int) (content map[string]interface{}, err error) {
	response, err := http.Get(fmt.Sprintf("%s/%d", authorHttpURL, authorId))

	if err != nil {
		return nil, newError(400, "Could not send request to the author service")
	}

	defer response.Body.Close()

	resultBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, newError(400, "Could not read response body")
	}

	if response.StatusCode != 200 {
		return nil, newErrorByte(400, resultBody)
	}

	var resultMap map[string]interface{}

	err = json.Unmarshal(resultBody, &resultMap)

	if err != nil {
		return nil, newError(400, "Could not unmarshal response body")
	}

	return resultMap, nil
}
