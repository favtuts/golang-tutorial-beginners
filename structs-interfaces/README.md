# Exploring structs and interfaces in Go
* https://tuts.heomi.net/exploring-structs-and-interfaces-in-go/

# Init a Go project

Locate the project directory
```bash
$ mkdir structs-interfaces
$ cd structs-interfaces
```

Then create the module
```bash
$ go mod init github.com/favtuts/structs-interfaces
```

Create main.go with the content:
```go
package main  // specifies the package name

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

# What are structs in Go?

Let’s say we have a blog post that we intend to publish. Using a struct type to represent the data fields would look like this:
```go
type blogPost struct {
  author  string    // field
  title   string    // field  
  postId  int       // field
}
```

Now, to instantiate or initialize the struct using a literal, we can do the following:
```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func NewBlogPost() *blogPost {
        return &blogPost{
                author: "Alexander",
                title:  "Learning structs and interfaces in Go",
                postId: 4555,
        }

}

func main() {
        var b blogPost // initialize the struct type

        fmt.Println(b) // print the zero value    

        newBlogPost := *NewBlogPost()
        fmt.Println(newBlogPost)

        // alternatively
        b = blogPost{ //
        author: "Alex",
        title: "Understand struct and interface types",
        postId: 12345,
        }

        fmt.Println(b)        

}

//output
{Alexander Learning structs and interfaces in Go 4555}
{  0}  // zero values of the struct type is shown
{Alex Understand struct and interface types 12345}
```

Run the code:
```bash
$ go run .

{  0}
{Alexander Learning structs and interfaces in Go 4555}
{Alex Understand struct and interface types 12345}
```

We can also use the dot, `.`, operator to access individual fields in the struct type after initializing them. Let’s see how we would do that with an example:
```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        var b blogPost // b is a type Alias for the BlogPost
        b.author= "Alex"
        b.title="understand structs and interface types"
        b.postId=12345

        fmt.Println(b)  

        b.author = "Chinedu"  // since everything is pass by value by default in Go, we can update this field after initializing - see pointer types later

        fmt.Println("Updated Author's name is: ", b.author)           
}
```

Run the code:
```bash
$ go run .

{Alex understand structs and interface types 12345}
Updated Author's name is:  Chinedu
```