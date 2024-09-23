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

	// queryRow(db)
	// queryRows(db)
	queryWithParameters(db)
}

func queryRow(db *sql.DB) {

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

func queryRows(db *sql.DB) {

	rows, err := db.Query("SELECT bird, description FROM birds limit 10")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}
	// create a slice of birds to hold our results
	birds := []Bird{}

	// iterate over the returned rows
	// we can go over to the next row by calling the `Next` method, which will
	// return `false` if there are no more rows
	for rows.Next() {
		bird := Bird{}
		// create an instance of `Bird` and write the result of the current row into it
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		// append the current instance to the slice of birds
		birds = append(birds, bird)
	}
	// print the length, and all the birds
	fmt.Printf("found %d birds: %+v", len(birds), birds)

}

func queryWithParameters(db *sql.DB) {

	birdName := "eagle"

	// For Postgres, parameters are specified using the "$" symbol, along with the index of
	// the param. Variables should be added as arguments in the same order
	// The sql library takes care of converting types from Go to SQL based on the driver
	row := db.QueryRow("SELECT bird, description FROM birds WHERE bird = $1 LIMIT $2", birdName, 1)

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
