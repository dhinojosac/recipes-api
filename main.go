// @title           Recipes API
// @description     his is a sample recipes API. You can find out more about the API at
// https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes		 http
// @host      localhost:8080
// @BasePath  /
// @version         1.0

// @consumes	   application/json
// @produces	   application/json

// @securityDefinitions.basic  BasicAuth
package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "recipes-api/docs"
	"recipes-api/handlers"
)

var recipesHandler *handlers.RecipesHandler

func init() {
	ctx := context.Background()

	// Mongo
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(os.Getenv(
		"MONGO_DATABASE")).Collection("recipes")

	// Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: "",
		DB:       0,
	})
	stcCmd := redisClient.Ping(ctx)
	if stcCmd.Err() != nil {
		log.Fatal(stcCmd.Err())
	}
	log.Println("Connected to Redis")

	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.GET("/recipes/:id", recipesHandler.GetRecipeByIDHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
