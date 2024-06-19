package main

import (
	"os"

	"golang-restaurant-management/database"

	middleware "golang-restaurant-management/middleware"
	routes "golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)


var foodCollection * mongo.Collection = database.OpenCollection(database.Client,"food")

// var foodCollection *mongo.Collection: This declares a variable foodCollection which is a pointer to a MongoDB collection.
// database.OpenCollection(database.Client,"food"): This calls a function OpenCollection from the database package, passing a Client and the collection name "food". This function likely returns a pointer to the MongoDB collection named "food".

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// router.Use(gin.Logger()): This adds a middleware to log the details of incoming HTTP requests.


	routes.UserRoutes(router)

	router.Use(middleware.Authentication())
// the line router.Use(middleware.Authentication()) adds an authentication middleware to secure all the routes that are defined after this middleware is applied. In the context of the provided code, it means that all the routes that are defined after this line will require authentication. Here's the breakdown:

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}