package main

import (
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"

	"github.com/VlasovArtem/distributed-system-example/authors/internal/config"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/handler/rpc"
	pb "github.com/VlasovArtem/distributed-system-example/grpc/authors"
)

func main() {
	logger, err := zap.NewDevelopment(zap.WithCaller(true))
	if err != nil {
		log.Fatal("error init logger", err)
	}
	zap.ReplaceGlobals(logger)

	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logger.Error("error process config", zap.Error(err))
	}
	logger.Sugar().Debugf("config: %+v", cfg)
	authorsService := service.New()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.RPC.TCPPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthorsServer(s, &rpc.Server{
		Service: authorsService,
	})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
