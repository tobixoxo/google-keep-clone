package main

import (
	"fmt"
	"time"


	"context"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc){
	defer cancel()

	defer func(){
		if err := client.Disconnect(ctx); err != nil{
			fmt.Println("application shutdown error: ", err)
		}
	}()
}

// returns mongodb client, context for timeout, cancelfunction and error
func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error){
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Hour)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error{
	if err := client.Ping(ctx, readpref.Primary()); err!= nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}