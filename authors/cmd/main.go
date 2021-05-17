package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"gitlab.lohika.com/dmiroshnichenko/distributed-comm-stubs/authors/internal/config"
	"gitlab.lohika.com/dmiroshnichenko/distributed-comm-stubs/authors/internal/handler/rest"
	"gitlab.lohika.com/dmiroshnichenko/distributed-comm-stubs/authors/internal/service"
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

	(&http.Server{
		Addr:    ":" + strconv.Itoa(cfg.HTTP.Port),
		Handler: rest.New(service.New()),
	}).ListenAndServe()
}
