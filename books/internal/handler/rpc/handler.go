package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/VlasovArtem/distributed-system-example/books/internal/service"
	"github.com/VlasovArtem/distributed-system-example/grpc/authors"
	pb "github.com/VlasovArtem/distributed-system-example/grpc/books"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type Server struct {
	pb.UnimplementedBooksServer
	Service *service.Service
}

func (s *Server) GetBooks(context.Context, *emptypb.Empty) (*pb.BooksResponse, error) {
	var books []*pb.Book
	for _, i2 := range s.Service.List() {
		books = append(books, &pb.Book{
			ID:          i2.ID,
			Title:       i2.Title,
			Description: i2.Description,
			AuthorID:    i2.AuthorID,
		})
	}
	return &pb.BooksResponse{
		Books: books,
	}, nil
}

func (s *Server) FindBook(_ context.Context, bookRequest *pb.FindBookRequest) (*pb.Book, error) {
	book, err := s.Service.GetByID(bookRequest.ID)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.Book{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		AuthorID:    book.AuthorID,
	}, nil
}

func (s *Server) AddBook(_ context.Context, book *pb.Book) (*pb.Book, error) {
	existsById := s.Service.ExistsById(book.ID)

	if existsById {
		return nil, errors.New(fmt.Sprintf("Book with id %s already exists.", existsById))
	}

	if err := updateAuthor(s.Service, book.AuthorID); err != nil {
		return nil, err
	}

	return book, nil
}

func updateAuthor(service *service.Service, authorId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	author, err := service.AuthorsClient.FindAuthor(ctx, &authors.FindAuthorRequest{
		ID: authorId,
	})

	if err != nil {
		return err
	}

	author.NumberOfBooks = author.NumberOfBooks + 1

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = service.AuthorsClient.UpdateAuthor(ctx, author)

	if err != nil {
		return err
	}

	return nil
}
