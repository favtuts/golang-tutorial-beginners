package models

import (
	"log"
	"time"

	"github.com/username/notes_api_layered/config"
)

type Note struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteParams struct {
	Title string
	Body  string
}

func (note *Note) Create(data NoteParams) (*Note, error) {
	var created_at = time.Now().UTC()
	var updated_at = time.Now().UTC()
	statement, _ := config.DB.Prepare("INSERT INTO notes (title, body, created_at, updated_at) VALUES (?, ?, ?, ?)")
	result, err := statement.Exec(data.Title, data.Body, created_at, updated_at)
	if err == nil {
		id, _ := result.LastInsertId()
		note.Id = int(id)
		note.Title = data.Title
		note.Body = data.Body
		note.CreatedAt = created_at
		note.UpdatedAt = updated_at
		return note, err
	}
	log.Println("Unable to create note", err.Error())
	return note, err
}

func (note *Note) GetAll() ([]Note, error) {
	rows, err := config.DB.Query("SELECT * FROM notes")
	allNotes := []Note{}
	if err == nil {
		for rows.Next() {
			var currentNote Note
			rows.Scan(
				&currentNote.Id,
				&currentNote.Title,
				&currentNote.Body,
				&currentNote.CreatedAt,
				&currentNote.UpdatedAt)
			allNotes = append(allNotes, currentNote)
		}
		return allNotes, err
	}
	return allNotes, err
}

func (note *Note) Fetch(id string) (*Note, error) {
	err := config.DB.QueryRow(
		"SELECT id, title, body, created_at, updated_at FROM notes WHERE id=?", id).Scan(
		&note.Id, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
	return note, err
}
