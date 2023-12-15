package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Note struct {
	NoteId string `json:"noteId"`
	Title string `json:"title"`
	Content string  `json:"content"`
}


var notes = []Note {
	{NoteId: "123", Title:"my holiday entry", Content: "it was super fun!"},
}

func createNote(c *gin.Context) {
	var newNote Note
	if err := c.BindJSON(&newNote); err != nil {
		fmt.Println("Error binding JSON:", err)
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func getNotes(c *gin.Context){
	c.IndentedJSON(http.StatusOK, notes)
}



func main() {
	router := gin.Default()
	
	router.POST("/createNote", createNote)
	router.GET("/getNotes", getNotes)

	router.Run("localhost:5000")
}