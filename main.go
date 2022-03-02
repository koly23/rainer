package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koly23/rainer/labels"
	"github.com/koly23/rainer/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

const apiPrefix = "/api/"

func search() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://10.32.179.223:27017"))
	if err != nil {
		logger.InfoE("cannot connect to mongodb", err)
	}
	collection := client.Database("test").Collection("books")
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		logger.InfoE("inert error", err)
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
	//r.GET("/labels", labels.All)
	route(http.MethodGet, "labels", labels.All, r)
	route(http.MethodPost, "labels", labels.Create, r)
	//r.POST("/labels", labels.Create)
	r.Run("localhost:9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func route(method string, path string, handler gin.HandlerFunc, r *gin.Engine) {
	path = apiPrefix + path
	r.Handle(method, path, handler)
}
