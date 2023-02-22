package controllers

//
//import (
//	"context"
//	"echo-mongo-api/configs"
//	"echo-mongo-api/models"
//	"echo-mongo-api/responses"
//	services "echo-mongo-api/services/user"
//	"fmt"
//	"github.com/labstack/echo/v4"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"net/http"
//	"time"
//)
//
//var transactionCollection *mongo.Collection = configs.GetCollection(configs.DB, "transactions")
//
//func CreateTransaction(c echo.Context) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	userId := c.Param("userId")
//	var transaction models.Transaction
//	var user models.User
//	defer cancel()
//
//	fmt.Println(transaction.UserId)
//
//	userObjId, _ := primitive.ObjectIDFromHex(userId)
//
//	err := services.UserCollection.FindOne(ctx, bson.M{"id": userObjId}).Decode(&user)
//
//	fmt.Println(err)
//
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, responses.UserResponse{
//			Status:  http.StatusInternalServerError,
//			Message: "error",
//			Data:    &echo.Map{"data": "erro"},
//		})
//	}
//
//	if err := c.Bind(&transaction); err != nil {
//		return c.JSON(http.StatusBadRequest, responses.UserResponse{
//			Status:  http.StatusBadRequest,
//			Message: "error",
//			Data:    &echo.Map{"data": err.Error()},
//		})
//	}
//
//	if validationErr := services.validate.Struct(&transaction); validationErr != nil {
//		return c.JSON(http.StatusBadRequest, responses.UserResponse{
//			Status:  http.StatusBadRequest,
//			Message: "error",
//			Data:    &echo.Map{"data": validationErr.Error()},
//		})
//	}
//
//	newTransaction := models.Transaction{
//		Value:     transaction.Value,
//		Category:  transaction.Category,
//		UserId:    transaction.UserId,
//		CreatedAt: time.Now(),
//	}
//
//	result, err := transactionCollection.InsertOne(ctx, newTransaction)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, responses.UserResponse{
//			Status:  http.StatusInternalServerError,
//			Message: "error",
//			Data:    &echo.Map{"data": err.Error()},
//		})
//	}
//
//	return c.JSON(http.StatusCreated, responses.UserResponse{
//		Status:  http.StatusCreated,
//		Message: "success",
//		Data:    &echo.Map{"data": result},
//	})
//
//}
