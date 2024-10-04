# Golang Constructors - Design Patterns For Initializing Variables
* https://tuts.heomi.net/golang-constructors-design-patterns-for-initializing-variables/

# Init a Go project

Locate the project directory
```bash
$ mkdir constructor-examples
$ cd constructor-examples
```

Then create the module
```bash
$ go mod init github.com/favtuts/constructor-examples
```

# Using Composite Literals

Composite literals are the most straight-forward way to initialize an object in Go.

Constructing a variable using composite literals looks like this:

```go
// creating a new struct instance
b := Book{}

// creating a pointer to a struct instance
bp := &Book{}

// creating an empty value
nothing := struct{}{}
```

Let's assume we have a Book struct defined:
```go
type Book struct {
  title string
  pages int
}
```

Now, we can initialize a new Book instance with its attributes set:
```go
b := Book{
  title: "Julius Caesar",
  pages: 322,
}
```

The advantage of using composite literals is that they have a *straightforward syntax* that’s easy to read.

However, we cannot set default values to each attribute. So, when we have structs that contain many default fields, we would need to repeat the default value for each instantiation.

For example, consider a `Pizza` struct, where we want six slices by default:
```go
type Pizza struct {
  slices int
  toppings []string
}

somePizza := Pizza{
  slices: 6,
  toppings: []string{"pepperoni"},
}

otherPizza := Pizza{
  slices: 6,
  toppings: []string{"onion", "pineapple"},
}
```

In this example, we have to keep setting the number of slices as 6 each time.

Additionally, when we have nullable attributes like `toppings`, it would be set to `nil` if not explicitly initialized. This can be error prone if your code assumes that each attribute is initialized.

Run the codes:
```bash
$ go run using_composite_litterals.go 
{ 0}
&{ 0}
{ 0}
{}
{Julius Caesar 322}
{6 [pepperoni]}
{6 [onion pineapple]}
```

# Custom Constructor Functions

Making your own constructor function can be useful if you need to set defaults or perform some initialization steps beforehand.

Let’s look at how we can construct our `Pizza` instance better by creating a `NewPizza` function:
```go
func NewPizza(toppings []string) () {
  if toppings == nil {
    toppings = []string{}
  }
  return Pizza{
    slices: 6,
    toppings: toppings,
  }
}
```

Run the code:
```bash
$ go run custom_constructor_functions.go 
{6 []}
{6 [pepperoni]}
{6 [onion pineapple]}
```

# Returning Errors from Constructors

While constructing your variable, we are sometimes dependent on other systems or libraries that may fail.

In these cases, it’s better to return an `error` along with the initialized value.

```bash
func NewRemotePizza(url string) (Pizza, error) {
  // toppings are received from a remote URL, which may fail
  toppings, err := getToppings(url)
  if err != nil {
    // if an error occurs, return the wrapped error along with an empty
    // Pizza instance
    return Pizza{}, fmt.Errorf("could not construct new Pizza: %v", err)
  }
  return Pizza{
    slices:   6,
    toppings: toppings,
  }, nil
}
```