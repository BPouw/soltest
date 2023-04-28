package controller

import (
	"context"
	"net/http"
	"time"
	"webshop/api/config"
	"webshop/api/model"
	"webshop/api/response"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var itemCollection *mongo.Collection = config.GetCollection(config.DB, "items")
var validate = validator.New()

func CreateItem(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var item model.Item
	defer cancel()

	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, response.ItemResponse{Status: http.StatusBadRequest, Message: "Error", Data: &echo.Map{"data": err.Error()}})
	}

	// use the validator library to validate required fields
	if validationErr := validate.Struct(&item); validationErr != nil {
		return c.JSON(http.StatusBadRequest, response.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newItem := model.Item{
		Id:          primitive.NewObjectID(),
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Stock:       item.Stock,
		Seller:      item.Seller,
	}

	result, err := itemCollection.InsertOne(ctx, newItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, response.ItemResponse{Status: http.StatusCreated, Message: "succes", Data: &echo.Map{"data": result}})

}

func GetItem(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	itemId := c.Param("itemId")
	var item model.Item
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(itemId)

	err := itemCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&item)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, response.ItemResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": err.Error()}})

}
