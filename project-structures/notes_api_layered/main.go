package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/username/notes_api_layered/config"
	"github.com/username/notes_api_layered/controllers"
	"github.com/username/notes_api_layered/migrations"
)

func main() {
	_, err := config.InitializeDB()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		// Run all migrations
		migrations.Run()

		router := gin.Default()

		var noteController controllers.NoteController
		router.GET("/notes", noteController.GetAllNotes)
		router.POST("/notes", noteController.CreateNewNote)
		router.GET("/notes/:note_id", noteController.GetSingleNote)
		router.Run(":8000")
	}
}
