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

	filter := bson.M{
		"id": objId,
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	updater := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "deleted_at",
					Value: time.Now(),
				},
			},
		},
	}

	result, err := UserCollection.UpdateOne(ctx, filter, updater)

	if err != nil {
		return err
	}

	if result.ModifiedCount < 1 {
		return errors.New("user with this ID not exists")
	}

	return nil
}
