# Using a PostgreSQL Database in Go (in Production) 
* https://tuts.heomi.net/using-a-postgresql-database-in-go-in-production/

# Init a Go project

Locate the project directory
```bash
$ mkdir structs-interfaces
$ cd structs-interfaces
```

Then create the module
```bash
$ go mod init github.com/favtuts/postgresql-database
```

Create main.go with the content:
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, Go")
}
```

Run the main function
```bash
$ go run .
Hello, Go
```

# Installing the Database Driver

We’ll be using the [pgx](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql#getting-started-with-pgx-through-databasesql) driver
```bash
$ go get github.com/jackc/pgx/v5
$ go get github.com/jackc/pgx/v5/pgxpool@v5.7.1
```

# Setup Postgres docker container

Run the container
```bash
$ docker pull postgres
$ docker run --name postgres-container -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres 
```

Log into the container:
```bash
$ docker exec -it postgres-container bash
$ docker exec -it postgres-container psql -U root
```

Then use `psql` to connect with `root` user:
```
# psql -U root
```

To exit from the Postgres server just type `\q`
```bash
root=# \q
```

List all databases
```bash
root=# \l
```

# Creating a Test Database

Create database and user:
```bash
$ create database bird_encyclopedia;
$ create user favtuts with encrypted password 'password';
$ grant all privileges on database bird_encyclopedia to favtuts;
$ GRANT ALL ON SCHEMA public TO favtuts;
$ ALTER DATABASE bird_encyclopedia OWNER TO favtuts;
$ exit
```

Connect inside container with `favtuts` user on the database `bird_encyclopedia`:
```bash
$ psql -U favtuts bird_encyclopedia
```

Connect remotely with `favtuts` user:
```bash
$ psql postgresql://favtuts:password@localhost:5432/bird_encyclopedia
```

List all tables of database `bird_encyclopedia`
```bash
$ \dt 

or run the query

$ SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';
```

Create table `birds` on database `bird_encyclopedia`:
```bash
CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  bird VARCHAR(256),
  description VARCHAR(1024)
);
```

Checking again
```bash
bird_encyclopedia=> \dt
        List of relations
 Schema | Name  | Type  |  Owner
--------+-------+-------+---------
 public | birds | table | favtuts
(1 row)
```

Insert sample data:
```bash
INSERT INTO birds (bird , description) VALUES 
('pigeon', 'common in cities'),
('eagle', 'bird of prey');
```

# Opening a Database Connection

We can now use the installed driver to open and verify a new connection with our database.
```go
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
}
```

The `db` variable in this example is an instance of [`*sql.DB`](https://pkg.go.dev/database/sql#DB) which represents the reference to our database instance

Although `pgx` provides its [own interface](https://github.com/jackc/pgx?tab=readme-ov-file#choosing-between-the-pgx-and-databasesql-interfaces) for interacting with PostgreSQL, we are utilizing the `pgx` adapter with the `database/sql` interface. This approach allows us to leverage the standard `database/sql` package’s features while still benefiting from some of the performance and advanced features of `pgx`.

If you’re running the database on your local machine, you can set the `DATABASE_URL` environment variable to `postgresql://localhost:5432/bird_encyclopedia` before running the code.

```bash
$ export DATABASE_URL="postgresql://favtuts:password@localhost:5432/bird_encyclopedia"
$ go run .

database is reachable
```

You can set the environment variable directly in the command
```bash
$ DATABASE_URL="postgresql://favtuts:password@localhost:5432/bird_encyclopedia" go run main.go

database is reachable
```

You can also set the environment variable to run debug in Visual Studio Code by creating the `launch.json` file
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "DATABASE_URL": "postgresql://favtuts:password@localhost:5432/bird_encyclopedia"
            }
        }
    ]
}
```

# Executing SQL Queries

We can use the `db.QueryRow` method when we require a single entry from our table (For example, fetching an entry based on its unique key).

First, let’s define a struct to represent the results of each query:
```go
type Bird struct {
	Species     string
	Description string
}
```

Let’s use the `QueryRow` method to fetch the first entry in our `birds` table:
```go
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
```

This will give us the following output:
```bash
found bird: {Species:pigeon Description:common in cities}
```

If we want to query multiple rows, we can use the `Query` method, which returns a [Rows](https://pkg.go.dev/database/sql#Rows) instance instead of a single row like the previous example.
```go
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
```

Output:
```bash
found 2 birds: [{Species:pigeon Description:common in cities} {Species:eagle Description:bird of prey}]
```

# Adding Query Parameters

We can use the `Query` and `QueryRow` methods to add variables in our code as query parameters. To illustrate, let’s add a `WHERE` clause to get information about the eagle:
```go
birdName := "eagle"
// For Postgres, parameters are specified using the "$" symbol, along with the index of
// the param. Variables should be added as arguments in the same order
// The sql library takes care of converting types from Go to SQL based on the driver
row := db.QueryRow("SELECT bird, description FROM birds WHERE bird = $1 LIMIT $2", birdName, 1)

// the code to scan the obtained row is the same as before
//...
```

Note: The symbols used for query params depends on the database you’re using. For example, we have to use the `?` symbol in MySQL instead of `$` which is specific to Postgres.

Output:
```bash
found bird: {Species:eagle Description:bird of prey}
```

# Executing Writes – INSERT, UPDATE, and DELETE

When we do writing to a database, it will returns the number of rows affected.

With the `sql` library, we can make use of the `Exec` method to execute write queries. Let’s see how we can use this to insert a new entry into the `birds` table:
```go
// sample data that we want to insert
newBird := Bird{
	Species:     "rooster",
	Description: "wakes you up in the morning",
}
// the `Exec` method returns a `Result` type instead of a `Row`
// we follow the same argument pattern to add query params
result, err := db.Exec("INSERT INTO birds (bird, description) VALUES ($1, $2)", newBird.Species, newBird.Description)
if err != nil {
	log.Fatalf("could not insert row: %v", err)
}

// the `Result` type has special methods like `RowsAffected` which returns the
// total number of affected rows reported by the database
// In this case, it will tell us the number of rows that were inserted using
// the above query
rowsAffected, err := result.RowsAffected()
if err != nil {
	log.Fatalf("could not get affected rows: %v", err)
}
// we can log how many rows were inserted
fmt.Println("inserted", rowsAffected, "rows")
```

Output:
```bash
inserted 1 rows
```