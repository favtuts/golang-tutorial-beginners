# Structuring your Golang app: Flat structure vs. layered architecture
* https://tuts.heomi.net/structuring-your-golang-app-flat-structure-vs-layered-architecture/

# Building a simple API using a flat structure

Create a new directory for this project by running:
```sh
mkdir notes_api_flat
```

Now, initialize the project:
```sh
go mod init github.com/username/notes_api_flat
```

We’ll use [SQLite3](https://www.sqlite.org/index.html) for storing the notes and [Gin](https://gopkg.in/gin-gonic/gin.v1) for routing.
```sh
go get github.com/mattn/go-sqlite3
go get github.com/gin-gonic/gin
```

Next, create the following files:

* `main.go`: entry point to the application
* `models.go`: manages access to the database
* `migration.go`: manages creating tables

After creating them, the folder structure should look like this:
```sh
notes_api_flat/
  go.mod
  go.sum
  go.main.go
  migration.go
  models.go
```

## Writing migration.go

Add the following to migration.go to create the table that will store our notes.
```go
package main
import (
  "database/sql"
  "log"
)
const notes = `
  CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(64) NOT NULL,
    body MEDIUMTEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
  )
`
func migrate(dbDriver *sql.DB) {
  statement, err := dbDriver.Prepare(notes)
  if err == nil {
    _, creationError := statement.Exec()
    if creationError == nil {
      log.Println("Table created successfully")
    } else {
      log.Println(creationError.Error())
    }
  } else {
    log.Println(err.Error())
  }
}
```

## Creating models.go

Add the following to `models.go`:
```go
package main
import (
  "log"
  "time"
)
type Note struct {
  Id        int       `json:"id"`
  Title     string    `json:"title"`
  Body      string    `json:"body"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
func (note *Note) create(data NoteParams) (*Note, error) {
  var created_at = time.Now().UTC()
  var updated_at = time.Now().UTC()
  statement, _ := DB.Prepare("INSERT INTO notes (title, body, created_at, updated_at) VALUES (?, ?, ?, ?)")
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
func (note *Note) getAll() ([]Note, error) {
  rows, err := DB.Query("SELECT * FROM notes")
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
  err := DB.QueryRow(
    "SELECT id, title, body, created_at, updated_at FROM notes WHERE id=?", id).Scan(
    &note.Id, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
  return note, err
}
```

## Completing the API in Go

The final piece remaining in the API is routing. Modify `main.go` to include the following code:
```go
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
```


## Running and Testing the API 

1. **Run the Application**: Start the Gin server and ensure that the API is running on the specified port.
```bash
$ go run .


2024/08/28 10:22:15 Table created successfully
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /notes                    --> main.getAllNotes (3 handlers)
[GIN-debug] POST   /notes                    --> main.createNewNote (3 handlers)
[GIN-debug] GET    /notes/:note_id           --> main.getSingleNote (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
```

2. **Test the API**: Use tools like Postman or curl to test the API endpoints. Ensure that all CRUD operations work as expected.

Get all notes
```bash
curl --location 'http://localhost:8000/notes'
```

Create new note
```bash
curl --location 'http://localhost:8000/notes' \
--header 'Content-Type: application/json' \
--data '{
    "title": "My Note 1",
    "body": "This is greating note"
}'
```

Get single note
```bash
curl --location 'http://localhost:8000/notes/1'
```

# Using a layered architecture (classic MVC structure) in Go

Let’s see what a layered architecture structure looks like:
```bash
layered_app/
  app/
    models/
      User.go         
    controllers/
      UserController.go
  config/
    app.go
  views/
    index.html
  public/
    images/
      logo.png
  main.go
  go.mod
  go.sum
```

Although layered architecture is not ideal for building simple libraries, it’s well suited for building APIs and other large applications.

# Updating the Go app with a layered architecture

Create a new folder called `notes_api_layered` and initialize a Go module in it by running the snippet below:
```bash
mkdir notes_api_layered
cd notes_api_layered
go mod init github.com/username/notes_api_layered
```

Install the required SQLite and Gin packages.
```bash
go get github.com/mattn/go-sqlite3
go get github.com/gin-gonic/gin
```

Now, update the project folder structure to look like this:
```bash
notes_api_layered/
  config/
    db.go
  controllers/
    note.go
  migrations/
    main.go
    note.go
  models/
    note.go
  go.mod
  go.sum
  main.go
```

# Completing the layered Go app with Go codes


## Database Configuration

Starting with the database configuration, add the following to `config/db.go`:
```go
package config
import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
func InitializeDB() (*sql.DB, error) {
  // Initialize connection to the database
  var err error
  DB, err = sql.Open("sqlite3", "./notesapi.db")
  return DB, err
}
```

We create a `DB` variable that will hold the connection to the database, as it’s not ideal for each model to have different instances of the database. 

**Note: starting a variable name or function name with capital letters means they should be exported.**

Then we declare and export an InitializeDB function, which opens the database and stores the database reference in the `DB` variable.


## Database Migration

We have two files in the migrations folder: `main.go` and `note.go`.
* `main.go` handles loading the queries to be performed, then performing them, 
* `note.go` contains SQL queries specific to the notes table

Now, add the following to `migrations/note.go`:
```go
package migrations
const Notes = `
CREATE TABLE IF NOT EXISTS notes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(64) NOT NULL,
  body MEDIUMTEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
)
`
```

Update `migrations/main.go` to include:
```go
package migrations
import (
  "database/sql"
  "log"
  "github.com/username/notes_api_layered/config"
)
func Run() {
  // Migrate notes
  migrate(config.DB, Notes)
  // Other migrations can be added here.
}
func migrate(dbDriver *sql.DB, query string) {
  statement, err := dbDriver.Prepare(query)
  if err == nil {
    _, creationError := statement.Exec()
    if creationError == nil {
      log.Println("Table created successfully")
    } else {
      log.Println(creationError.Error())
    }
  } else {
    log.Println(err.Error())
  }
}
```

## Update the models

After applying these changes, `models/note.go` should look like this:
```go
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
```

We’ve declared a new package, `models`, and we imported the config from `github.com/username/notes_api_layered/config`. With that, we have access to the `DB` that would have been initialized once the `InitializeDB` function gets called.


## Change the controller

The controller will look like this:
```bash
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
```

From the above snippet, we converted the functions to methods so that they can be accessed via `NoteController.Create` instead of `controller.Create`. This is to account for the possibility of having multiple controllers, which would be the case for most modern applications.


## Update main application

Finally, we update `main.go` to reflect the refactoring:
```bash
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
```

Following refactoring, we have a lean main package that imports the required packages: `config`, `controllers`, and `models`. Then, it initializes the database by calling `config.InitializeDB()`.

Now we can move on to routing. The routes should be updated to use the newly created note controller for handling requests.


## Running and Testing the API 

1. **Run the Application**: Start the Gin server and ensure that the API is running on the specified port.
```bash
$ go run .


$ go run .
2024/08/28 14:14:56 Table created successfully
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /notes                    --> github.com/username/notes_api_layered/controllers.(*NoteController).GetAllNotes-fm (3 handlers)
[GIN-debug] POST   /notes                    --> github.com/username/notes_api_layered/controllers.(*NoteController).CreateNewNote-fm (3 handlers)
[GIN-debug] GET    /notes/:note_id           --> github.com/username/notes_api_layered/controllers.(*NoteController).GetSingleNote-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
```

2. **Test the API**: Use tools like Postman or curl to test the API endpoints. Ensure that all CRUD operations work as expected.

Get all notes
```bash
curl --location 'http://localhost:8000/notes'
```

Create new note
```bash
curl --location 'http://localhost:8000/notes' \
--header 'Content-Type: application/json' \
--data '{
    "title": "My Note 1",
    "body": "This is greating note"
}'
```

Get single note
```bash
curl --location 'http://localhost:8000/notes/1'
```