package services

import (
	"context"
	"echo-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FindAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User

	filter := bson.M{
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	results, err := UserCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleUser models.User
		if err := results.Decode(&singleUser); err != nil {
			return nil, err
		}
		users = append(users, singleUser)
	}

	return users, nil
}
