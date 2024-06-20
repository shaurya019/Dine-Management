package controller

import (
	"context"
	"fmt"
	"golang-restaurant-management/database"
	"golang-restaurent-management/models"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()


func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {}
}


func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food id")
		var food models.Food

		err := foodCollection.FindOne(ctx,bson.M{"food_id":foodId}).Decode(&food)
		// FindOne(ctx, bson.M{"food_id": foodId}): This method is used to find a single document in the foodCollection collection that matches the filter criteria. The filter is specified using a BSON map (bson.M), where food_id is the field you are matching against the value of foodId
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}


func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// c.BindJSON(&food): This function attempts to bind the JSON payload from the HTTP request body to the food struct. Here, c is the context of the Gin framework, which encapsulates the request and response.


		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price,2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}


func UpdateFood() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel() // Ensure context cancellation is called regardless of code path
        
        var menu models.Menu
        var food models.Food

        foodId := c.Param("food_id")

        if err := c.BindJSON(&food); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var updateObj primitive.D

        if food.Name != nil {
            updateObj = append(updateObj, bson.E{"name", food.Name})
        }

        if food.Price != nil {
            updateObj = append(updateObj, bson.E{"price", food.Price})
        }

        if food.Food_image != nil {
            updateObj = append(updateObj, bson.E{"food_image", food.Food_image})
        }

        if food.Menu_id != nil {
            err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
            if err != nil {
                msg := fmt.Sprintf("message:Menu was not found")
                c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
                return
            }
            // Include menu_id in the update object
            updateObj = append(updateObj, bson.E{"menu_id", food.Menu_id})
        }

        food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        updateObj = append(updateObj, bson.E{"updated_at", food.Updated_at})

        upsert := true
        filter := bson.M{"food_id": foodId}

        opt := options.UpdateOptions{
            Upsert: &upsert,
        }

        result, err := foodCollection.UpdateOne(
            ctx,
            filter,
            bson.D{
                {"$set", updateObj},
            },
            &opt,
        )

        if err != nil {
            msg := fmt.Sprintf("food item update failed")
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
        }
        c.JSON(http.StatusOK, result)
    }
}

