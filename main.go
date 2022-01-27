package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func search() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://10.32.179.223:27017"))
	if err != nil {
		fmt.Println("cannot connect to mongodb")
		fmt.Println(err)
	}
	collection := client.Database("test").Collection("books")
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		fmt.Println("error")
	}
	id := res.InsertedID
	fmt.Println("id ", id)
}

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	search()
	//r.POST("/search", search)
	r.Run("localhost:18080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
