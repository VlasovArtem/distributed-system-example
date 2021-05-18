package rpc

import (
	"context"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	pb "github.com/VlasovArtem/distributed-system-example/grpc/authors"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type Server struct {
	pb.UnimplementedAuthorsServer
	Service *service.Service
}

func (s *Server) GetAuthors(context.Context, *emptypb.Empty) (*pb.AuthorsResponse, error)  {
	var authors []*pb.Author
	for _, i2 := range s.Service.List() {
		authors = append(authors, &pb.Author{
			ID: i2.ID,
			FirstName: i2.FirstName,
			LastName: i2.LastName,
		})
	}
	return &pb.AuthorsResponse{
		Authors: authors,
	}, nil
}

func (s *Server) FindAuthor(ctx context.Context, authorRequest *pb.FindAuthorRequest) (*pb.Author, error)  {
	author, err := s.Service.GetByID(authorRequest.ID)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.Author{
		ID: author.ID,
		FirstName: author.FirstName,
		LastName: author.LastName,
	}, nil
}