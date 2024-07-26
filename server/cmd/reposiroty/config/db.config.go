package db_config

import (
	"billAPI/internals/models"
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = "billSplitter"

var collections = map[string]string{
	"users": "users",
	"logs":  "logs",
}

func InitUserCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(database).Collection(collections["users"])

}

func InitializeDB() (*mongo.Client, *models.AppError) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if err != nil {
		return nil, &models.AppError{
			Message: "failed to connect to MongoDB",
			Code:    500,
			Err:     err,
		}
	}

	UserCollection := InitUserCollection(client)

	// check if the user with _id equal to 0 exists
	result := UserCollection.FindOne(context.TODO(), bson.M{"_id": 0})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			// if there is an error that means that the user has not been found so we can create a new one
			user := models.User{UserName: "test", Password: "asñdlkfjñsdal", Id: 0}

			_, err := UserCollection.InsertOne(context.TODO(), user)
			if err != nil {
				return nil, &models.AppError{
					Message: "failed to create user with _id 0",
					Code:    500,
					Err:     err,
				}
			}
		} else {
			// unexpected error while querying MongoDB
			return nil, &models.AppError{
				Message: "error querying MongoDB for user with _id 0",
				Code:    500,
				Err:     err,
			}
		}
	} else {
		slog.Info("User with _id 0 exists")
	}

	slog.Info("Database initialized")
	return client, nil
}
