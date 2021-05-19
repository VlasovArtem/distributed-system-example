package main

import (
	"github.com/VlasovArtem/distributed-system-example/books/internal/handler/mq"
	"github.com/VlasovArtem/distributed-system-example/books/internal/handler/rpc"
	"log"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/VlasovArtem/distributed-system-example/books/internal/config"
	"github.com/VlasovArtem/distributed-system-example/books/internal/handler/rest"
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
	var msq *mq.MessageQueueService

	if cfg.MQ.URL != "" {
		msq = mq.StartMessageQueueConnection(&cfg)

		defer msq.CloseMessageQueueService()
	}

	booksService := service.New(msq)

	if cfg.RPC.Enabled {
		rpc.StartRPCServer(booksService, &cfg)
	} else {
		(&http.Server{
			Addr:    ":" + strconv.Itoa(cfg.HTTP.Port),
			Handler: rest.New(booksService),
		}).ListenAndServe()
	}
}
