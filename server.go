package main

import (
	"github.com/gin-gonic/gin"
)

type Note struct {
	NoteId string `bson:"noteId"`
	Title string `bson:"title"`
	Content string  `bson:"content"`
}


func main() {

	client, ctx, cancel, err := connect("mongodb://localhost:27017")

	if err!= nil{
		panic(err)
	}
	defer close(client, ctx, cancel)

	router := gin.Default()

	router.POST("/createNote", func(c *gin.Context){
		createNote(c, client)
	})

	router.GET("/getNotes", func(c *gin.Context) {
		getNotes(c, client, ctx)
	})
	router.DELETE("/deleteNote/")

	router.Run("localhost:5000")
}