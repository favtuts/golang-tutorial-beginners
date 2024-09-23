#  A Guide On SQL Database Transactions In Go
* https://tuts.heomi.net/a-guide-on-sql-database-transactions-in-go/

# Init a Go project

Locate the project directory
```bash
$ mkdir sql-transactions
$ cd sql-transactions
```

Then create the module
```bash
$ go mod init github.com/favtuts/sql-transactions
```

Install the dependencies
```bash
$ go get github.com/lib/pq
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
$ create database pet_shop;
$ create user favtuts with encrypted password 'password';
$ grant all privileges on database pet_shop to favtuts;
$ GRANT ALL ON SCHEMA public TO favtuts;
$ ALTER DATABASE pet_shop OWNER TO favtuts;
$ exit
```

Connect inside container with `favtuts` user on the database `pet_shop`:
```bash
$ psql -U favtuts pet_shop
```

Connect remotely with `favtuts` user:
```bash
$ psql postgresql://favtuts:password@localhost:5432/pet_shop
```

List all tables of database `pet_shop`
```bash
$ \dt 

or run the query

$ SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';
```

Create table `pets` and `food` on database `pet_shop` as in [migrations.sql](./migrations.sql) file
```bash
CREATE TABLE pets (id SERIAL PRIMARY KEY,name TEXT,species TEXT);
CREATE TABLE food (id SERIAL PRIMARY KEY, name TEXT, quantity INT);
```

Checking again
```bash
pet_shop=> \dt
        List of relations
 Schema | Name | Type  |  Owner
--------+------+-------+---------
 public | food | table | favtuts
 public | pets | table | favtuts
(2 rows)
```

# Basic Transactions


Assume we execute 2 queries:
```sql
INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat');
INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5);
```

Now lets think about what happens if the first query succeeds, but the second query fails: you now have data which shows that two new pets are adopted, but no food has been bought.

* We cannot treat this as a success, since one of the queries failed
* We cannot treat this as a failure, since the first query passed. If we consider this a failure and retry, we would have to insert `Fido` and `Albert` for a second time.


![without-tx](./images/without-tx.png)


To avoid situations like this, we want both the queries to pass or fail together. This is where SQL transactions come in. If you see the example below, it’s just a slight modification from executing the queries normally:
```sql
BEGIN;
INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat');
INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5);
END;
```

* The `BEGIN` statement starts a new transaction
* The `END` statement commits the above transactions atomically

![with-tx](./images/with-tx.png)

We will now make use of the [`pq`](github.com/lib/pq) library, along with the `database/sql` standard library, and implement the two insert queries we discussed in the previous example:
```go
package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Create a new connection to our database
	connStr := "postgresql://favtuts:password@localhost:5432/pet_shop?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// `tx` is an instance of `*sql.Tx` through which we can execute our queries

	// Here, the query is executed on the transaction instance, and not applied to the database yet
	_, err = tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return
	}

	// The next query is handled similarly
	_, err = tx.ExecContext(ctx, "INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5)")
	if err != nil {
		tx.Rollback()
		return
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
```

Here, we can commit or rollback the transaction from our Go code itself.

# Read-and-Update Transactions

In some cases, we want to update data in a table depending on the previous values stored in some other table. Transactions can help us do this atomically as well.

For example, consider the case where we want to add cat food to our inventory, and the quantity that we want to add is proportional to the number of cats that we have.

In this case, we would need to find the count of the total number of cats first, and then update the quantity of cat food.

```go
package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Initialize a connection, and begin a transaction like before    
	connStr := "postgresql://favtuts:password@localhost:5432/pet_shop?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
	if err != nil {
		tx.Rollback()
		return
	}

	// Run a query to get a count of all cats
	row := tx.QueryRow("SELECT count(*) FROM pets WHERE species='cat'")
	var catCount int
	// Store the count in the `catCount` variable
	err = row.Scan(&catCount)
	if err != nil {
		tx.Rollback()
		return
	}

	// Now update the food table, increasing the quantity of cat food by 10x the number of cats
	_, err = tx.ExecContext(ctx, "UPDATE food SET quantity=quantity+$1 WHERE name='Cat Food'", 10*catCount)
	if err != nil {
		tx.Rollback()
		return
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
```

Here, we can allow our application to decide how much cat-food to add to our inventory (which in this case is 10x the number of cats).

It’s important to note why the read query is executed within the transaction: any read query outside the transaction doesn’t consider the values of an uncommitted transaction. This means that if our read query was outside the transaction, we would not consider the pets added in the first insert query.

![read-in-tx](./images/read-in-tx.png)


# Build and run the main.go file:

Build the application
```bash
go build -o transaction-example
```

Run the application
```bash
./transaction-example
```

You can run the application directly
```bash
$ go run .
```