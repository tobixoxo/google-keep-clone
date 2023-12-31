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

func deleteNote(c *gin.Context, client *mongo.Client, ctx context.Context){
	noteId := c.Param("noteId")
	deletectx := context.Background()
	filter := bson.D{{"noteId",noteId}}
	coll := client.Database("google-keep-clone-db").Collection("keep-notes")
	result, err := coll.DeleteOne(deletectx, filter)
	if err != nil {
		fmt.Println("error deleting note: ", err)
		c.Status(400)
	}
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)
	c.Status(200)
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

func updateNote(c *gin.Context, client *mongo.Client, ctx context.Context){
	// capture id, title and content
	var updatedNote Note
	if err := c.BindJSON(&updatedNote); err != nil {
		fmt.Println("Error binding JSON:", err)
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	noteId := updatedNote.NoteId
	updatectx := context.Background()
	filter := bson.D{{"noteId",noteId}}
	fmt.Println(updatedNote.NoteId)
	fmt.Println(updatedNote.Title)
	fmt.Println(updatedNote.Content)
	update := bson.D{{"$set", bson.D{{"title", updatedNote.Title}, {"content", updatedNote.Content}}}}
	coll := client.Database("google-keep-clone-db").Collection("keep-notes")

	result, err := coll.UpdateOne(updatectx, filter, update)
	if err != nil {
		fmt.Println("error updating note: ", err)
		return 
	}
	if result.ModifiedCount == 0 {
		fmt.Println("no Note updated");
	}
	if result.MatchedCount == 0 {
		fmt.Println("no Note matched")
	}
	// fmt.Println(result.ModifiedCount)
	c.IndentedJSON(http.StatusCreated, result)


	// get note by id

	// update the note 

	// return 
}