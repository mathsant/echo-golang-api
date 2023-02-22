package services

import (
	"context"
	"echo-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func UpdateUser(id string, dataToUpdate bson.M) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := UserCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": dataToUpdate})

	if err != nil {
		return nil, err
	}

	var updatedUser models.User
	if result.ModifiedCount == 1 {
		err := UserCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return nil, err
		}
	}

	return &updatedUser, nil
}
