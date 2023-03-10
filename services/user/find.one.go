package services

import (
	"context"
	"echo-mongo-api/configs"
	"echo-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func FindOneUser(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"id": objId,
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	err := UserCollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
