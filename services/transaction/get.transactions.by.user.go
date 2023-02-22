package services

import (
	"context"
	"echo-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetTransactionsByUser(userId string) (*[]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var transactions []models.Transaction

	filter := bson.M{
		"userid": userId,
		"deleted_at": bson.M{
			"$exists": false,
		},
	}

	results, err := TransactionCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleTransaction models.Transaction
		if err := results.Decode(&singleTransaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, singleTransaction)
	}

	return &transactions, nil
}
