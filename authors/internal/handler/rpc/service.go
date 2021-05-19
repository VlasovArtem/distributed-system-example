package rpc

import (
	"github.com/VlasovArtem/distributed-system-example/authors/internal/config"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	pb "github.com/VlasovArtem/distributed-system-example/grpc/authors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

type Server struct {
	pb.UnimplementedAuthorsServer
	Service *service.Service
}

func StartRPCServer(s *service.Service, cfg *config.Config) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.RPC.TCPPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthorsServer(server, &Server{
		Service: s,
	})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
