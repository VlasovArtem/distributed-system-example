package main

import (
	pb "github.com/VlasovArtem/distributed-system-example/grpc/books"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/VlasovArtem/distributed-system-example/books/internal/config"
	"github.com/VlasovArtem/distributed-system-example/books/internal/handler/rest"
	"github.com/VlasovArtem/distributed-system-example/books/internal/handler/rpc"
	"github.com/VlasovArtem/distributed-system-example/books/internal/service"
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

	if cfg.RPC.Enabled {
		lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.RPC.TCPPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterBooksServer(s, &rpc.Server{
			Service: authorsService,
		})
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	} else if cfg.HTTP.Enabled {
		(&http.Server{
			Addr:    ":" + strconv.Itoa(cfg.HTTP.Port),
			Handler: rest.New(authorsService),
		}).ListenAndServe()
	}
}
