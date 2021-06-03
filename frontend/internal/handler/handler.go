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
	"log"
	"net/http"
	"time"
)

type result struct {
	ID              string
	Title           string
	Description     string
	AuthorID        string
	AuthorFirstName string
	AuthorLastName  string
}

type grpcService struct {
	authorsClient authors.AuthorsClient
	booksClient   books.BooksClient
}

func New(config *config.Config) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	grpcService := grpcService{
		authorsClient: openAuthorsRPCConnection(config.Authors.RPC),
		booksClient:   openBooksRPCConnection(config.Books.RPC),
	}

	router.HandleFunc("/api/v1/dashboard", GetDashboard(&grpcService)).Methods("GET")
	router.HandleFunc("/api/v1/info", Info(config)).Methods("GET")

	return router
}

func Info(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(config)
	}
}

func GetDashboard(service *grpcService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var results []result

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

		var authorsResult = make(map[string]*authors.Author)

		for _, author := range authorsResponse.Authors {
			authorsResult[author.ID] = author
		}

		for _, book := range booksResponse.Books {
			results = append(results, result{
				ID: book.ID,
				Title: book.Title,
				Description: book.Description,
				AuthorID: book.AuthorID,
				AuthorFirstName: authorsResult[book.AuthorID].FirstName,
				AuthorLastName: authorsResult[book.AuthorID].LastName,
			})
		}

		json.NewEncoder(w).Encode(results)
	}
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
