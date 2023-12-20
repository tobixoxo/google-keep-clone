package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
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

	router.Use(static.Serve("/", static.LocalFile("../my-react-app/build", true)))

	router.POST("/notes", func(c *gin.Context){
		createNote(c, client)
	})

	router.GET("/notes", func(c *gin.Context) {
		getNotes(c, client, ctx)
	})

	router.DELETE("/notes/:noteId",
	func(c *gin.Context) {
		deleteNote(c, client, ctx)
	})

	router.PUT("/notes", func(c * gin.Context){
		updateNote(c, client, ctx)
	})
	

	router.Run("localhost:5000")
}