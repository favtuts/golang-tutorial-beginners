# Golang Interfaces Explained
* https://tuts.heomi.net/golang-interfaces-explained/


# Setup project module
```bash
$ go mod init golanginterfaces
```

# What is an interface in Go?

An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.

One example of an interface type from the standard library is the [fmt.Stringer](https://pkg.go.dev/fmt/#Stringer) interface, which looks like this:
```go
type Stringer interface {
    String() string
}
```

The following `Book` type satisfies the interface because it has a `String() string` method:
```go
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}
```

The following `Count` type also satisfies the `fmt.Stringer` interface
```go
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}
```

Wherever you see declaration in Go (such as a variable, function parameter or struct field) which has an interface type, you can use an object of any type so long as it satisfies the interface.
```go
func WriteLog(s fmt.Stringer) {
    log.Print(s.String())
}
```

Because this `WriteLog()` function uses the `fmt.Stringer` interface type in its parameter declaration, we can pass in any object that satisfies the `fmt.Stringer` interface. For example, we could pass either of the `Book` and `Count` types that we made earlier to the `WriteLog()` method, and the code would work OK.

The full example code:
```go
package main

import (
    "fmt"
    "strconv"
    "log"
)

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
    log.Print(s.String())
}

func main() {
    // Initialize a Count object and pass it to WriteLog().
    book := Book{"Alice in Wonderland", "Lewis Carrol"}
    WriteLog(book)

    // Initialize a Count object and pass it to WriteLog().
    count := Count(3)
    WriteLog(count)
}
```

Test the `fmt.Stringer` interface demo:
```bash
$ go run .
2024/08/29 15:09:54 Book: Alice in Wonderland - Lewis Carrol
2024/08/29 15:09:54 3
```

# Reducing boilerplate code

We have a Customer struct containing some data about a customer. In one part of our codebase we want to write the customer information to a [bytes.Buffer](https://pkg.go.dev/bytes/#Buffer), and in another part of our codebase we want to write the customer information to an [os.File](https://pkg.go.dev/os/#File) on disk. But in both cases, we want to serialize the customer struct to JSON first.

The first thing you need to know is that Go has an [io.Writer](https://pkg.go.dev/io/#Writer) interface type which looks like this:
```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

And we can leverage the fact that both `bytes.Buffer` and the `os.File` type satisfy this interface, due to them having the `bytes.Buffer.Write()` and `os.File.Write()` methods respectively.

Let’s take a look at a simple implementation:
```go
package main

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "os"
)

// Create a Customer type
type Customer struct {
    Name string
    Age  int
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to JSON, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.
func (c *Customer) WriteJSON(w io.Writer) error {
    js, err := json.Marshal(c)
    if err != nil {
        return err
    }

    _, err = w.Write(js)
    return err
}

func main() {
    // Initialize a customer struct.
    c := &Customer{Name: "Alice", Age: 21}

    // We can then call the WriteJSON method using a buffer...
    var buf bytes.Buffer
    err := c.WriteJSON(&buf)
    if err != nil {
        log.Fatal(err)
    }

    // Or using a file.
    f, err := os.Create("/tmp/customer")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()


    err = c.WriteJSON(f)
    if err != nil {
        log.Fatal(err)
    }
}
```

we can create the `Customer.WriteJSON()` method once, and we can call that method any time that we want to write to something that satisfies the `io.Writer` interface.

We can test the code:
```bash
$ go run .
$ cat /tmp/customer
{"Name":"Alice","Age":21}
```

# Unit testing and mocking

Let’s say you run a shop, and you store information about the number of customers and sales in a PostgreSQL database. You want to write some code that calculates the sales rate (i.e. sales per customer) for the past 24 hours, rounded to 2 decimal places.

First install the Go postgres driver
```bash
go get github.com/lib/pq
```

A minimal implementation of the code for that could look something like this:

**File: shop.go**

```go
package main

import (
    "fmt"
    "log"
    "time"
    "database/sql"
    _ "github.com/lib/pq"
)

type ShopDB struct {
    *sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
    var count int
    err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
    return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
    var count int
    err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
    return count, err
}

func main() {
    db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    shopDB := &ShopDB{db}
    sr, err := calculateSalesRate(shopDB)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf(sr)
}

func calculateSalesRate(sdb *ShopDB) (string, error) {
    since := time.Now().Add(-24 * time.Hour)

    sales, err := sdb.CountSales(since)
    if err != nil {
        return "", err
    }

    customers, err := sdb.CountCustomers(since)
    if err != nil {
        return "", err
    }

    rate := float64(sales) / float64(customers)
    return fmt.Sprintf("%.2f", rate), nil
}
```

Now, what if we want to create a unit test for the `calculateSalesRate()` function to make sure that the math logic in it is working correctly?

A solution here is to create our own interface type which describes the `CountSales()` and `CountCustomers()` methods that the `calculateSalesRate()` function relies on. Then we can update the signature of `calculateSalesRate()` to use this custom interface type as a parameter, instead of the concrete `*ShopDB` type.

Re-update the **File: shop.go**:
```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Create our own custom ShopModel interface. Notice that it is perfectly
// fine for an interface to describe multiple methods, and that it should
// describe input parameter types as well as return value types.
type ShopModel interface {
	CountCustomers(time.Time) (int, error)
	CountSales(time.Time) (int, error)
}

// The ShopDB type satisfies our new custom ShopModel interface, because it
// has the two necessary methods -- CountCustomers() and CountSales().
type ShopDB struct {
	*sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}
	sr, err := calculateSalesRate(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sr)
}

// Swap this to use the ShopModel interface type as the parameter, instead of the
// concrete *ShopDB type.
func calculateSalesRate(sm ShopModel) (string, error) {
	since := time.Now().Add(-24 * time.Hour)

	sales, err := sm.CountSales(since)
	if err != nil {
		return "", err
	}

	customers, err := sm.CountCustomers(since)
	if err != nil {
		return "", err
	}

	rate := float64(sales) / float64(customers)
	return fmt.Sprintf("%.2f", rate), nil
}
```

With that done, it’s straightforward for us to create a mock which satisfies our `ShopModel` interface. We can then use that mock during unit tests to test that the math logic in our `calculateSalesRate()` function works correctly. Like so:

**File: shop_test.go**

```go
package main

import (
    "testing"
    "time"
)

type MockShopDB struct{}

func (m *MockShopDB) CountCustomers(_ time.Time) (int, error) {
    return 1000, nil
}

func (m *MockShopDB) CountSales(_ time.Time) (int, error) {
    return 333, nil
}

func TestCalculateSalesRate(t *testing.T) {
    // Initialize the mock.
    m := &MockShopDB{}
    // Pass the mock to the calculateSalesRate() function.
    sr, err := calculateSalesRate(m)
    if err != nil {
        t.Fatal(err)
    }

    // Check that the return value is as expected, based on the mocked
    // inputs.
    exp := "0.33"
    if sr != exp {
        t.Fatalf("got %v; expected %v", sr, exp)
    }
}
```

You could run that test now, everything should work fine.
```bash
$ go test -run "^TestCalculateSalesRate$"

PASS
ok      golanginterfaces        0.002s
```

# What is the empty interface?

An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.

The empty interface type: `interface{}`. The empty interface type essentially describes no methods. It has no rules. And because of that, it follows that any and every object satisfies the empty interface. The empty interface type interface{} is kind of like a wildcard, you can declare the empty interface to use an object of any type.

Example:
```go
package main

import "fmt"


func main() {
    person := make(map[string]interface{}, 0)

    person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

    fmt.Printf("%+v", person)
}
```

In this code snippet we initialize a `person` map, which uses the `string` type for keys and the empty interface type `interface{}` for values. We’ve assigned three different types as the map values (a `string`, `int` and `float32`) — and that’s OK. Because objects of any and every type satisfy the empty interface, the code will work just fine.

You can test it as follow:
```bash
s$ go run .
map[age:21 height:167.64 name:Alice]
```

There’s an important thing to point out when it comes to retrieving and using a value from this map. Let’s say that we want to get the "age" value and increment it by 1. If you write something like the following code, it will fail to compile:
```go
person["age"] = person["age"] + 1
```

And you’ll get the following error message:
```bash
invalid operation: person["age"] + 1 (mismatched types interface {} and int)
```

To get around this this, you need to type assert the value back to an int before using it
```go
age, ok := person["age"].(int)
if !ok {
    log.Fatal("could not assert value to int")
    return
}

person["age"] = age + 1
```

So when should you use the empty interface type in your own code?

The answer is *probably not that often*. If you find yourself reaching for it, pause and consider whether using `interface{}` is really the right option. As a general rule it’s clearer, safer and more performant to use concrete types — or non-empty interface types — instead. In the code snippet above, it would have been more appropriate to define a `Person` struct with relevant typed fields similar to this:
```go
type Person struct {
    Name   string
    Age    int
    Height float32
}
```

But that said, the empty interface is useful in situations where you need to accept and work with unpredictable or user-defined types.

# The any identifier

Go 1.18 introduced a new [predeclared identifier](https://tip.golang.org/ref/spec#Predeclared_identifiers) called [any](https://pkg.go.dev/builtin#any), which is an alias for the empty interface `interface{}`. So writing `map[string]any` in your code is exactly the same as writing `map[string]interface{}` in terms of it’s behavior. In most modern Go codebases, you’ll normally see `any` being used rather than `interface{}`.