package handler

import (
	"encoding/json"
	"github.com/VlasovArtem/distributed-system-example/frontend/internal/config"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type result struct {
	ID              string
	Title           string
	Description     string
	AuthorID        string
	AuthorFirstName string
	AuthorLastName  string
}

func New(config *config.Config) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/dashboard", GetDashboardByHTTP(config.Books.URL, config.Authors.URL))
	router.HandleFunc("/api/v1/info", Info(config))

	return router
}

func Info(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(config)
	}
}

func GetDashboardByHTTP(booksDest, authorsDest string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booksResult []map[string]interface{}
		booksData, err := getData(booksDest)

		if err == nil {
			err = json.Unmarshal(booksData, &booksResult)
		}

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Could retrieve data from books service"))
		}

		var authorsResult []map[string]interface{}

		booksData, err = getData(authorsDest)

		if err == nil {
			err = json.Unmarshal(booksData, &authorsResult)
		}

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Could retrieve data from authors service"))
		}

		newAuthorsResult := make(map[string]map[string]interface{})

		for _, author := range authorsResult {
			newAuthorsResult[author["ID"].(string)] = author
		}

		var results []result

		for _, book := range booksResult {
			authorID := book["AuthorID"].(string)
			results = append(results, result{
				ID:              book["ID"].(string),
				Title:           book["Title"].(string),
				Description:     book["Description"].(string),
				AuthorID:        authorID,
				AuthorFirstName: newAuthorsResult[authorID]["FirstName"].(string),
				AuthorLastName:  newAuthorsResult[authorID]["LastName"].(string),
			})
		}

	}
}

func getData(url string) (data []byte, err error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return result, nil
}
