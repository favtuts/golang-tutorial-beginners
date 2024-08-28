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
