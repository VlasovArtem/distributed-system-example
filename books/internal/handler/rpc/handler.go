package rpc

import (
	"context"
	"github.com/VlasovArtem/distributed-system-example/books/internal/model"
	"github.com/VlasovArtem/distributed-system-example/books/internal/service"
	pb "github.com/VlasovArtem/distributed-system-example/grpc/books"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type Server struct {
	pb.UnimplementedBooksServer
	Service *service.Service
}

func (s *Server) GetBooks(context.Context, *emptypb.Empty) (*pb.BooksResponse, error) {
	var books []*pb.Book
	for _, i2 := range s.Service.List() {
		books = append(books, &pb.Book{
			ID:       i2.ID,
			Title:    i2.Title,
			Pages:    uint32(i2.Pages),
			AuthorID: uint32(i2.AuthorID),
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
		ID:       book.ID,
		Title:    book.Title,
		Pages:    uint32(book.Pages),
		AuthorID: uint32(book.AuthorID),
	}, nil
}

func (s *Server) AddBookAndAuthor(_ context.Context, in *pb.BookAndAuthor) (*emptypb.Empty, error) {
	bookAndAuthor := model.BookAndAuthor{
		Book: model.Book{
			ID:       in.ID,
			Title:    in.Title,
			Pages:    int(in.Pages),
			AuthorID: int(in.AuthorID),
		},
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	s.Service.Add(bookAndAuthor)

	return &emptypb.Empty{}, nil
}
