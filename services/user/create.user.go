package services

import (
	"context"
	"echo-mongo-api/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var validate = validator.New()

func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if validationErr := validate.Struct(&user); validationErr != nil {
		return nil, validationErr
	}

	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Title:    user.Title,
		Location: user.Location,
	}

	result, err := UserCollection.InsertOne(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return result, nil
}
