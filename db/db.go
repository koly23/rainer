package db

import (
	"context"
	"github.com/koly23/rainer/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const database = "rainer"
const Labels = "labels"

//const databaseUrl = "mongodb://10.32.179.223:27017"
const databaseUrl = "mongodb://127.0.0.1:27017"
const timeout = 10 * time.Second

type Db struct {
	client *mongo.Client
}

func NewDb() Db {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		logger.InfoE("connect db error", err)
	}
	return Db{
		client: client,
	}
}

func (db Db) All(collection string, page int, size int) []bson.D {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := db.client.Database(database).Collection(collection).Find(ctx, bson.D{})
	if err != nil {
		logger.InfoE("db find error", err)
	}
	var results []bson.D
	if err = cursor.All(ctx, &results); err != nil {
		logger.InfoE("db find all error", err)
	}
	logger.InfoA("find all labels", results)
	return results
}

func (db Db) Create(collection string, content interface{}) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := db.client.Database(database).Collection(collection).InsertOne(ctx, content)
	if err != nil {
		logger.InfoE("db insert error", err)
	}
	return result.InsertedID
}
