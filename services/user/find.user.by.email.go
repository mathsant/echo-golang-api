package services

import (
	"context"
	"echo-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FindUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	filter := bson.M{
		"email": email,
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
