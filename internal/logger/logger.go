package logger

import (
	"github.com/developwithayush/go-todo-app/internal/config"
	"go.uber.org/zap"
)


type Logger = *zap.Logger

func NewLogger(cfg *config.Config) Logger {
	var log *zap.Logger
	var err error

	if cfg.Env ==  "prod" || cfg.Env == "production" {
		log, err = zap.NewProduction()
	} else {
		log, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}

	return log
}
 


func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
