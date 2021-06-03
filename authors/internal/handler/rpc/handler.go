package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	pb "github.com/VlasovArtem/distributed-system-example/grpc/authors"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type Server struct {
	pb.UnimplementedAuthorsServer
	Service *service.Service
}

func (s *Server) GetAuthors(context.Context, *emptypb.Empty) (*pb.AuthorsResponse, error) {
	var authors []*pb.Author
	for _, author := range s.Service.List() {
		authors = append(authors, convertAuthorToGrpcAuthor(&author))
	}
	return &pb.AuthorsResponse{
		Authors: authors,
	}, nil
}

func (s *Server) FindAuthor(_ context.Context, authorRequest *pb.FindAuthorRequest) (*pb.Author, error) {
	author, err := s.Service.GetByID(authorRequest.ID)
	if err != nil {
		log.Fatal(err)
	}

	return convertAuthorToGrpcAuthor(&author), nil
}

func (s *Server) AddAuthor(_ context.Context, author *pb.Author) (*pb.Author, error) {
	if s.Service.ExistsById(author.ID) {
		return nil, errors.New(fmt.Sprintf("Author with id %s already exists.", author.ID))
	}

	s.Service.Save(*convertGrpcAuthorToAuthor(author))

	return author, nil
}
func (s *Server) UpdateAuthor(_ context.Context, author *pb.Author) (*pb.Author, error) {
	if !s.Service.ExistsById(author.ID) {
		return nil, errors.New(fmt.Sprintf("Author with id %s is not exists.", author.ID))
	}

	s.Service.Save(*convertGrpcAuthorToAuthor(author))

	return author, nil
}

func convertAuthorToGrpcAuthor(author *service.Author) *pb.Author {
	return &pb.Author{
		ID:            author.ID,
		FirstName:     author.FirstName,
		LastName:      author.LastName,
		Age:           uint32(author.Age),
		Biography:     author.Biography,
		NumberOfBooks: uint32(author.NumberOfBooks),
	}
}

func convertGrpcAuthorToAuthor(author *pb.Author) *service.Author {
	return &service.Author{
		ID:            author.ID,
		FirstName:     author.FirstName,
		LastName:      author.LastName,
		Age:           uint8(author.Age),
		Biography:     author.Biography,
		NumberOfBooks: uint(author.NumberOfBooks),
	}
}