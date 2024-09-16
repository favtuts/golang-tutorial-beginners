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

# Golang nested or embedded struct fields

We can nest the `Author` struct in the `blogPost` struct like this:
```go
package main

import "fmt"

type Author struct {
  firstName, lastName, Biography string
  photoId    int
}

type blogPost struct {
  author  Author // nested struct field
  title   string
  postId  int 
  published  bool  
}

func main() {
        b := new(blogPost)

        fmt.Println(b)

        b.author.firstName= "Alex"
        b.author.lastName= "Nnakwue"
        b.author.Biography = "I am a lazy engineer"
        b.author.photoId = 234333
        b.published=true
        b.title= "understand interface and struct type in Go"
        b.postId= 12345

        fmt.Println(*b)        

}

// output

&{{   0}  0 false}  // again default values
{{Alex Nnakwue I am a lazy engineer 234333} understand interface and struct type in Go 12345 true}
```

In Go, there is a concept of promoted fields for nested struct types. In this case, we can directly access struct types defined in an embedded struct without going deeper, that is, doing `b.author.firstName`. Let’s see how we can achieve this:
```go
package main

import "fmt"

type Author struct {
  firstName, lastName, Biography string
  photoId    int
}

type BlogPost struct {
  Author  // directly passing the Author struct as a field - also called an anonymous field orembedded type 
  title   string
  postId  int 
  published  bool  
}

func main() {
        b := BlogPost{
        Author: Author{"Alex", "Nnakwue", "I am a lazy engineer", 234333},
        title:"understand interface and struct type in Go",
        published:true,
        postId: 12345,
        }

        fmt.Println(b.firstName) // remember the firstName field is present on the Author struct?
        fmt.Println(b)        

}

//output
Alex
{{Alex Nnakwue I am a lazy engineer 234333} understand interface and struct type in Go 12345 true}
```

# What are method sets in Golang?

Methods in Go are special kinds of functions with a receiver.

A method set of a type, `T`, that consists of all methods declared with receiver types, `T`. Note that the receiver is specified via an extra parameter preceding the method name.

In Go, we can create a type with a behavior by defining a method on that type. In essence, a method set is a list of methods that a type must have to implement an interface. Let’s look at an example:

```go
// BlogPost struct with fields defined
type BlogPost struct {
  author  string
  title   string
  postId  int  
}

// Create a BlogPost type called (under) Technology
type Technology BlogPost
```

Methods can also be defined on other named types:
```go
// write a method that publishes a blogPost - accepts the Technology type as a pointer receiver
func (t *Technology) Publish() {
    fmt.Printf("The title on %s has been published by %s, with postId %d\n" , t.title, t.author, t.postId)
}

// alternatively similar to the above, if we choose not to define a new type 
func (b *BlogPost) Publish() {
    fmt.Printf("The title on %s has been published by %s, with postId %d\n" , t.title, b.author, b.postId)
}

// Create an instance of the type
t := Technology{"Alex","understand structs and interface types",12345}

// Publish the BlogPost -- This method can only be called on the Technology type
t.Publish()

// output
The title on understand structs and interface types has been published by Alex, with postId 12345
```


# What is a Golang interface?

As we mentioned in the last section, method sets add behavior to one or more types. However, interface types define one or more method sets.

A type, therefore, is said to implement an interface by implementing its methods. In that light, interfaces enable us to compose custom types that have a common behavior.

Method sets are basically method lists that a type must have for that type to implement that interface.

In Go, interfaces are implicit. This means that if every method belonging to the method set of an interface type is implemented by a type, then that type is said to implement the interface. To declare an interface:
```go
type Publisher interface {
    publish()  error
}
```

In the `publish()` interface method we set above, if a type (for example, a struct) implements the method, then we can say the type implements the interface. Let’s define a method that accepts a struct type `blogpost` below:
```go
func (b blogPost) publish() error {
   fmt.Println("The title has been published by ", b.author)
   return nil
}
```

Now to implement the interface:
```go
package main

import "fmt"

// interface definition
type Publisher interface {
     Publish()  error
}

type blogPost struct {
  author  string
  title   string
  postId  int  
}

// method with a value receiver
func (b blogPost) Publish() error {
   fmt. Printf("The title on %s has been published by %s, with postId %d\n" , b.title, b.author, b.postId)
   return nil
}

 func test(){

  b := blogPost{"Alex","understanding structs and interface types",12345}

  fmt.Println(b.Publish())

   d := &b   // pointer receiver for the struct type

   b.author = "Chinedu"


   fmt.Println(d.Publish())

}


func main() {

        var p Publisher

        fmt.Println(p)

        p = blogPost{"Alex","understanding structs and interface types",12345}

        fmt.Println(p.Publish())

        test()  // call the test function 

}

//output
<nil>
The title on understanding structs and interface types has been published by Alex, with postId 12345
<nil>
The title on understanding structs and interface types has been published by Alex, with postId 12345
<nil>
The title on understanding structs and interface types has been published by Chinedu, with postId 12345
<nil>
```

We can also alias interface types like this:
```go
type publishPost Publisher  // alias to the interface defined above - mostly suited for third-party interfaces
```


# Convert an interface to a struct in Golang

From Effective Go, to cast an interface to a struct, we can make use of the syntax notation below:
```go
v = x.(T)
```

Here, `x` is the interface type and `T` is the actual concrete type. In essence, `T` must implement the interface type of `x`.

To check for correctness and avoid a type mismatch, we can go further and make use of the syntax notation below:
```go
v, ok = x.(T)
```
In this case, the value of `ok` is true if the assertion holds. 

Let’s see a trivial example of using type assertions to work with both structs and interfaces below:
```go
package main

import "fmt"

type blogPost struct {
        Data interface{}
        postId int
}

func NewBlogPostStruct() interface{} {
        return &blogPost{postId: 1234, Data: "Alexander"}
}

func main() {
        blogPost := NewBlogPostStruct().(*blogPost)
        fmt.Println(blogPost.Data)
}
//returns
Alexander
```