package db

import (
	"context"
	"gol/the-basics/dev/do"
	"log"
	"time"

	"gol/the-basics/dev/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	CreateAuthUser(authUser do.CreateAuthUserRequest) bool
}

type mongoDatabase struct {
	client *mongo.Client
}

//TODO: Currently have an IP Address whitelist. Validate that all is good with it.
func NewDatabase(globalConfig *config.GlobalConfig) Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	log.Println(globalConfig.DatabaseUrl)
	clientOptions := options.Client().ApplyURI(globalConfig.DatabaseUrl).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return mongoDatabase{client: client}
}

func (this mongoDatabase) CreateAuthUser(authUser do.CreateAuthUserRequest) bool {
	return true //TODO
}
