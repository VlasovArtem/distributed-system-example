package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type author struct {
	ID        string
	FirstName string
	LastName  string
}

type book struct {
	ID       string
	Title    string
	Pages    int
	AuthorID int
}

type result struct {
	Books   []book
	Authors []author
}

func New(booksDest, authorsDest string) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/dashboard", GetDashboard(booksDest, authorsDest))

	return router
}

func GetDashboard(booksDest, authorsDest string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result result

		err := json.Unmarshal(getData(booksDest), &result.Books)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(getData(authorsDest), &result.Authors)

		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(result)
	}
}

func getData(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
