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

Further, we can use the short literal notation to instantiate a struct type without using field names, as shown below:
```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        b := blogPost{"Alex", "understand struct and interface type", 12345}
        fmt.Println(b)        

}
```

Note that with the approach above, we must always pass the field values in the same order in which they are declared in the struct type. Also, all the fields must be initialized.

Finally, if we have a struct type to use only once inside a function, we can define them inline, as shown below:
```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {

        // inline struct init
        b := struct {
          author  string
          title   string
          postId  int  
         }{
          author: "Alex",
          title:"understand struct and interface type",
          postId: 12345,
        }

        fmt.Println(b)           
}
```

Note that we can also initialize struct types with the `new` keyword. Example: 
```go
b := new(blogPost)
```

In that case, we can do the following:
```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        b := new(blogPost)

        fmt.Println(b) // zero value

        b.author= "Alex"
        b.title= "understand interface and struct type in Go"
        b.postId= 12345

        fmt.Println(*b)   // dereference the pointer     

}

//output
&{  0}
{Alex understand interface and struct type in Go 12345}
```


Run the code:
```bash
$ go run .
Hello, Go
&{  0}
{Alex understand interface and struct type in Go 12345}
```

Note that as we can see from the output, by using the `new` keyword, we allocate storage for the variable, `b` , which then initializes the zero values of our struct fields — in this case `(author="", title="", postId=0)`.

This then returns a pointer type, `*b`, containing the address of the above variables in memory.


# Golang pointer to a struct

We can initialize struct using pointers:
```go
package main

import "fmt"

type blogPost struct {
  author  string
  title   string
  postId  int  
}

func main() {
        b := &blogPost{
                author:"Alex",
                title: "understand structs and interface types",
                postId: 12345,
                }

        fmt.Println(*b)   // dereference the pointer value 

       fmt.Println("Author's name", b.author) // in this case Go would handle the dereferencing on our behalf
}
```

Run the code:
```bash
$ go run .

{Alex understand structs and interface types 12345}
Author's name Alex
```