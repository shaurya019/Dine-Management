package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// returns a pointer (*mongo.Client) to a MongoDB client.
func DBinstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Print(MongoDb)

	client,err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// context.Background(): This function returns a background context. Contexts in Go are used to manage cancellation signals and deadlines across API boundaries. The background context is typically used as a base context for creating other contexts.

	// context.WithTimeout(...): This function creates a child context (ctx) with a specified timeout duration (10*time.Second). This means that the ctx context will automatically be canceled (its Done channel will be closed) after 10 seconds. This is useful to ensure that operations do not hang indefinitely if they exceed a reasonable time limit.
	
	// ctx, cancel: context.WithTimeout returns two values:
	
	// ctx: This is the context object that carries the deadline and cancellation signals.
	// cancel: This is a function that, when called, cancels the context and releases its resources. It should be called when the operations associated with the context are complete or no longer needed.
	// 2. defer cancel()
	// defer: In Go, defer is used to delay the execution of a function until the surrounding function returns. Here, cancel() will be called when the surrounding function (likely the function where these lines are placed) exits, either normally or due to an error.
	
	// cancel(): This function cancels the ctx context, closing its Done channel. This cancellation will propagate to any operations that are listening on this context. It's essential to cancel the context to prevent resource leaks and unnecessary computation if the operation exceeds its deadline.
	
	// 3. err = client.Connect(ctx)
	// client.Connect(ctx): This method initiates a connection to the MongoDB server using the provided context (ctx). It takes into account the timeout specified in the ctx.
	
	// err = ...: This line assigns any error returned by client.Connect(ctx) to the err variable. If there's an error connecting to MongoDB within the specified timeout period, it will be captured here.

	err = client.Connect(ctx);
	// Attempts to establish a connection to the MongoDB server using the provided context (ctx). This operation respects the timeout set in the ctx.

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
