package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/notes_api_layered/models"
)

type NoteController struct{}

func (_ *NoteController) CreateNewNote(c *gin.Context) {
	var params models.NoteParams
	var note models.Note
	err := c.BindJSON(&params)
	if err == nil {
		_, creationError := note.Create(params)
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
func (_ *NoteController) GetAllNotes(c *gin.Context) {
	var note models.Note
	notes, err := note.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "All Notes",
			"notes":   notes,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
func (_ *NoteController) GetSingleNote(c *gin.Context) {
	var note models.Note
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
