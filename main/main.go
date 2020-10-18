package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"runtime/debug"
	"time"
)

func main() {
	var err error
	var timeout = time.Millisecond * 1000 * 30
	err = testMongoDB("mongodb://127.0.0.1:27017", timeout)
	if err != nil {
		panic(string(debug.Stack()))
	}

}

func testNetworkOverload() (err error) {
	NetworkOve
}

func testMongoDB(address string, timeOut time.Duration) (err error) {
	var mongoClient *mongo.Client
	var ctx context.Context
	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), timeOut)
	err = mongoClient.Connect(ctx)
	if err != nil {
		return
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), timeOut)
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	defer cancel()

	err = mongoClient.Disconnect(ctx)

	return
}
