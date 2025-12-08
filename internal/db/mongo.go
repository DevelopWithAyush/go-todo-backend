package db

import (
	"context"
	"time"

	"github.com/developwithayush/go-todo-app/internal/config"
	"github.com/developwithayush/go-todo-app/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
	Users  *mongo.Collection
	Todos  *mongo.Collection
)

func InitMongo(cfg *config.Config, logr logger.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	db := client.Database(cfg.MongoDB)

	Client = client
	DB = db
	Users = DB.Collection("users")
	Todos = DB.Collection("todos")

	logr.Info("Connected to MongoDB", logger.Field("uri", cfg.MongoURI),
		logger.Field("database", cfg.MongoDB),
	)

	return nil
}
