package web

import (
	"fmt"

	"github.com/dilly3/blood-donor/internal"
	"go.uber.org/zap"
)

type Config struct {
	Serv   internal.ICandidateServ
	Logger *zap.Logger
}

func NewConfig(serv internal.ICandidateServ) *Config {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Error from zap, =>", err)
	}
	return &Config{
		Serv:   serv,
		Logger: logger,
	}
}
