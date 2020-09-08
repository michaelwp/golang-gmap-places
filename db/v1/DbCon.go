package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func DbCon(dbName string) (db *mongo.Database, status string, err error){
	uri := os.Getenv("MONGODB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	//errHandler.ErrHandler("Error connect to mongodb", err)

	err = client.Ping(ctx, nil)
	//errHandler.ErrHandler("Error connect to mongodb", err)

	status = "Connected to MongoDB!"

	return client.Database(dbName), status, err
}
