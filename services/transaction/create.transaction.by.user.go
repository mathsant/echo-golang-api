package services

import (
	"context"
	"echo-mongo-api/configs"
	"echo-mongo-api/models"
	services "echo-mongo-api/services/user"
	"echo-mongo-api/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var validate = validator.New()

var TransactionCollection *mongo.Collection = configs.GetCollection(configs.DB, "transactions")

func CreateTransactionByUser(transactionParam models.Transaction, userId string) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := services.FindOneUser(userId)

	if err != nil {
		return nil, err
	}

	if validationErr := validate.Struct(&transactionParam); validationErr != nil {
		return nil, validationErr
	}

	if user.Budget < transactionParam.Value {
		return nil, errors.New("you dont have budget for this transaction")
	}

	newTransaction := models.Transaction{
		Value:     transactionParam.Value,
		Category:  transactionParam.Category,
		UserId:    userId,
		CreatedAt: time.Now(),
	}

	userUpdatedWithNewBudget := bson.M{"budget": utils.ToFixed(user.Budget-transactionParam.Value, 2)}

	_, err = services.UpdateUser(userId, userUpdatedWithNewBudget)

	if err != nil {
		return nil, err
	}

	result, err := TransactionCollection.InsertOne(ctx, newTransaction)

	if err != nil {
		return nil, err
	}

	return result, nil

}
