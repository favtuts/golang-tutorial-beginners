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