package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := UserCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		return errors.New("user with this ID not exists")
	}

	return nil
}
