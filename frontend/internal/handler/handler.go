package handler

import (
	"context"
	"encoding/json"
	"github.com/VlasovArtem/distributed-system-example/frontend/internal/config"
	"github.com/VlasovArtem/distributed-system-example/grpc/authors"
	"github.com/VlasovArtem/distributed-system-example/grpc/books"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net/http"
	"time"
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

type grpcService struct {
	authorsClient authors.AuthorsClient
	booksClient   books.BooksClient
}

func New(config *config.Config) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	var handlerFunc http.HandlerFunc

	if config.RPCEnabled {
		grpcService := grpcService{
			authorsClient: openAuthorsRPCConnection(config.Authors.RPC),
			booksClient:   openBooksRPCConnection(config.Books.RPC),
		}

		handlerFunc = GetDashboardByRPC(&grpcService)


	} else {
		handlerFunc = GetDashboardByHTTP(config.Books.URL, config.Authors.URL)
	}

	if handlerFunc == nil {
		log.Fatal("Handler func not provided")
	}

	router.HandleFunc("/api/v1/dashboard", handlerFunc)
	router.HandleFunc("/api/v1/info", Info(config))

	return router
}

func Info(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.RPCEnabled {
			w.Write([]byte("I am working with RPC"))
		} else if config.HTTPEnabled {
			w.Write([]byte("I am working with HTTP"))
		}
		json.NewEncoder(w).Encode(config)
	}
}

func GetDashboardByHTTP(booksDest, authorsDest string) http.HandlerFunc {
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

func GetDashboardByRPC(service *grpcService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result result

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		booksResponse, err := service.booksClient.GetBooks(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not retrieve books: %v", err)
		}

		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		authorsResponse, err := service.authorsClient.GetAuthors(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not retrieve authors: %v", err)
		}

		for _, book := range booksResponse.Books {
			result.Books = append(result.Books, convertBook(book))
		}

		for _, author := range authorsResponse.Authors {
			result.Authors = append(result.Authors, convertAuthor(author))
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

func openAuthorsRPCConnection(address string) authors.AuthorsClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return authors.NewAuthorsClient(conn)
}

func openBooksRPCConnection(address string) books.BooksClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return books.NewBooksClient(conn)
}

func convertBook(b *books.Book) (book book) {
	book.ID = b.ID
	book.Title = b.Title
	book.Pages = int(b.Pages)
	book.AuthorID = int(b.AuthorID)
	return
}

func convertAuthor(a *authors.Author) (author author) {
	author.ID = a.ID
	author.FirstName = a.FirstName
	author.LastName = a.LastName
	return
}