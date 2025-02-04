package database

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	clientInstance *mongo.Client
	clientError    error
	mongoOnce      sync.Once
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().
			ApplyURI("mongodb://localhost:27017").
			SetAppName("dashboard")

		clientInstance, clientError = mongo.Connect(clientOptions)
		if clientError != nil {
			log.Fatalf("Ошибка подключения к MongoDB: %s", clientError)
		}

		if err := clientInstance.Ping(ctx, nil); err != nil {
			log.Fatalf("Ошибка пингования MongoDB: %s", err)
		}

		log.Println("Успешное подключение к MongoDB")
	})

	return clientInstance, clientError
}
