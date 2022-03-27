package infrastructure

import (
	"auth/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	config config.MongoConfig
	logger *Logger
	db     *mongo.Database
}

func NewMongoClient(config config.MongoConfig, logger *Logger) *MongoClient {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Url))

	if err != nil {
		logger.Fatalf("[Mongo] error while create client: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		logger.Fatalf("[Mongo] error while open connection: %s", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logger.Fatalf("[Mongo] error while ping: %s", err)
	}

	logger.Infof("[Mongo] connected")

	return &MongoClient{
		config: config,
		logger: logger,
		db:     client.Database(config.DBName),
	}
}

func (mongo *MongoClient) Destroy(ctx context.Context) {
	mongo.db.Client().Disconnect(ctx)
}

func (mongo *MongoClient) DB() *mongo.Database {
	return mongo.db
}
