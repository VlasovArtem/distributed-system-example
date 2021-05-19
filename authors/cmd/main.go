package main

import (
	"github.com/VlasovArtem/distributed-system-example/authors/internal/handler/mq"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/handler/rest"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/handler/rpc"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"

	"github.com/VlasovArtem/distributed-system-example/authors/internal/config"
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

	if cfg.MQ.URL != "" {
		messageQueueConnection := mq.StartMessageQueueConnection(&cfg)

		defer messageQueueConnection.CloseMessageQueueService()

		messageQueueConnection.StartMessageQueueConsumer(authorsService)
	}

	if cfg.RPC.Enabled {
		rpc.StartRPCServer(authorsService, &cfg)
	} else {
		(&http.Server{
			Addr:    ":" + strconv.Itoa(cfg.HTTP.Port),
			Handler: rest.New(authorsService),
		}).ListenAndServe()
	}
}
