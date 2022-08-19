package db

import (
	"context"
	"time"

	"gol/the-basics/dev/config"
	"gol/the-basics/dev/do"
	// "go.mongodb.org/mongo-driver/mongo"
)

type mongoDatabase struct {
	// client  *mongo.Client
	context context.Context
}

const DatabaseName = "ServerlessInstance0"

//TODO: Currently have an IP Address whitelist. Validate that all is good with it.
func NewDatabase(globalConfig *config.GlobalConfig) IDatabase {
	// serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	// clientOptions := options.Client().ApplyURI(globalConfig.DatabaseUrl).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return mongoDatabase{client: client, context: ctx}
	return mongoDatabase{context: ctx}
}

func (this mongoDatabase) CreateAuthUser(username string, encryptedPassword string) (*do.CreateAuthUserResponse, error) {
	// _, err := this.client.Database(DatabaseName).Collection(collections.Users).InsertOne(
	// 	this.context,
	// 	do.CreateAuthUserRequest{
	// 		Username: username,
	// 		Password: encryptedPassword,
	// 	},
	// )

	// if err != nil {
	// 	return &do.CreateAuthUserResponse{}, fmt.Errorf("IDatabase.CreateAuthUser: error inserting new user %w", err)
	// }

	return &do.CreateAuthUserResponse{
		Id:       "tbd", //TODOs
		Username: username,
		Password: encryptedPassword,
	}, nil
}
