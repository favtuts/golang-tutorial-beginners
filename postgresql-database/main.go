// file: main.go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	// Importing pgx v5 for PostgreSQL database operations. The pgx package is used
	// directly for database connection and operations, replacing the standard database/sql package.
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	// The `sql.Open` function opens a new `*sql.DB` instance. We specify the driver name
	// and the URI for our database. Here, we're using a Postgres URI from an environment variable
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method with a context. If no error is returned, we can assume a successful connection
	if err := db.PingContext(context.Background()); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")

	queryRows(db)
}

func queryRows(db *sql.DB) {

	// `QueryRow` always returns a single row from the database
	row := db.QueryRow("SELECT bird, description FROM birds LIMIT 1")
	// Create a new `Bird` instance to hold our query results
	bird := Bird{}
	// the retrieved columns in our row are written to the provided addresses
	// the arguments should be in the same order as the columns defined in
	// our query
	if err := row.Scan(&bird.Species, &bird.Description); err != nil {
		log.Fatalf("could not scan row: %v", err)
	}
	fmt.Printf("found bird: %+v\n", bird)

}
