package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Create this to store instance to SQL
var DB *sql.DB

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./notesapi.db")
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		// Create all the tables
		migrate(DB)
		router := gin.Default()
		router.GET("/notes", getAllNotes)
		router.POST("/notes", createNewNote)
		router.GET("/notes/:note_id", getSingleNote)
		router.Run(":8000")
	}
}

type NoteParams struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createNewNote(c *gin.Context) {
	var params NoteParams
	var note Note
	err := c.BindJSON(&params)
	if err == nil {
		_, creationError := note.create(params)
		if creationError == nil {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Note created successfully",
				"note":    note,
			})
		} else {
			c.String(http.StatusInternalServerError, creationError.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
func getAllNotes(c *gin.Context) {
	var note Note
	notes, err := note.getAll()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "All Notes",
			"notes":   notes,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
func getSingleNote(c *gin.Context) {
	var note Note
	id := c.Param("note_id")
	_, err := note.Fetch(id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Single Note",
			"note":    note,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
