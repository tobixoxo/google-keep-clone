package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func createNote(c *gin.Context, client *mongo.Client) {
	var newNote Note
	if err := c.BindJSON(&newNote); err != nil {
		fmt.Println("Error binding JSON:", err)
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// notes = append(notes, newNote)
	readContext := context.Background()
	coll := client.Database("google-keep-clone-db").Collection("keep-notes")
	coll.InsertOne(readContext, newNote)
	
	c.IndentedJSON(http.StatusCreated, newNote)
}

func getNotes(c *gin.Context,client *mongo.Client,  ctx context.Context){
	readctx := context.Background()
	filter := bson.D{{}} 
	coll := client.Database("google-keep-clone-db").Collection("keep-notes")
	cursor, err := coll.Find(readctx, filter )
	if err != nil {
		fmt.Println("error querying notes: ", err)
		return 
	}


	var result []Note
	if err := cursor.All(readctx, &result); err != nil {
		fmt.Println("error converting results: ", err)
		return 
	}

	c.IndentedJSON(http.StatusOK, result)
}