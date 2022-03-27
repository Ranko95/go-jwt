package main

import (
	"auth/config"
	"auth/internal/infrastructure"
	"context"
)

func main() {
	config := config.GetConfig()

	logger := infrastructure.NewLogger()

	mongoClient := infrastructure.NewMongoClient(config.MongoConfig, logger)

	defer mongoClient.Destroy(context.Background())

	httpServer := infrastructure.NewHttpServer(config.HTTPServerConfig, logger)

	httpServer.Run()
}
