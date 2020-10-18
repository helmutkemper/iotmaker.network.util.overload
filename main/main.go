package main

import (
	"context"
	"encoding/hex"
	"fmt"
	overload "github.com/helmutkemper/iotmaker.network.util.overload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"runtime/debug"
	"time"
)

func main() {
	var err error
	var timeout = time.Millisecond * 1000 * 30
	var delayMin = time.Millisecond * 500
	var delayMax = time.Millisecond * 5000

	err = testMongoDB("mongodb://127.0.0.1:27017", timeout)
	if err != nil {
		panic(string(debug.Stack()))
	}

	err = testNetworkOverload(delayMin, delayMax, "127.0.0.1:27016", "127.0.0.1:27017")
	if err != nil {
		panic(string(debug.Stack()))
	}

	err = testNetworkOverloaded("mongodb://127.0.0.1:27016", timeout)
	if err != nil {
		panic(string(debug.Stack()))
	}
}

func binaryDump(inData []byte, inLength int, direction overload.Direction) (outData []byte, outLength int, err error) {
	outData = inData
	outLength = inLength

	fmt.Printf("%v:\n", direction)
	fmt.Printf("%v\n", hex.Dump(inData[:inLength]))

	return
}

func testNetworkOverload(min, max time.Duration, inAddress, outAddress string) (err error) {
	var over = &overload.NetworkOverload{
		ConnectionInterface: &overload.TCPConnection{},
	}
	err = over.SetAddress(overload.KTypeNetworkTcp, inAddress, outAddress)
	if err != nil {
		return
	}

	over.ParserAppendTo(binaryDump)
	over.SetDelay(min, max)

	go func() {
		err = over.Listen()
		if err != nil {
			panic(string(debug.Stack()))
		}
	}()

	return
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

func testNetworkOverloaded(address string, timeout time.Duration) (err error) {
	start := time.Now()

	var mongoClient *mongo.Client
	var cancel context.CancelFunc
	var ctx context.Context

	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return
	}

	err = mongoClient.Connect(ctx)
	if err != nil {
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	type Trainer struct {
		Name string
		Age  int
		City string
	}
	collection := mongoClient.Database("test").Collection("trainers")
	ash := Trainer{"Ash", 10, "Pallet Town"}
	_, err = collection.InsertOne(context.TODO(), ash)
	if err != nil {
		return
	}
	fmt.Printf("fim\n")
	duration := time.Since(start)
	fmt.Printf("Duration: %v\n\n", duration)

	return
}
