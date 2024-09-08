# Type assertions vs. type conversions in Go
* https://tuts.heomi.net/type-assertions-vs-type-conversions-in-go/

# Init a Go project

Locate the project directory
```bash
$ mkdir type-assertion-conversion
$ cd type-assertion-conversion
```

Then create the module
```bash
$ go mod init github.com/favtuts/type-assertion-conversion
```

# Type assertion in Go

In Go, the syntax for type assertions is `t := i.(type)`. Here is a snippet of a full type assertion operation:
```go
// type-lessons.go

package main

func main() {
  var i interface{} = "a string"
  t := i.(string) // "a string"
}
```

The type assertion operation consists of three main elements:

* `i`, which is the variable whose type we are asserting. This variable must be defined as an interface
* `type`, which is the type we are asserting our variable is (such as string, int, float64, etc)
* `t`, which is a variable that stores the value of our variable i, if our type assertion is correct

It should be further noted that when type assertions are unsuccessful, it throws an error referred to as a “panic” in Go.

The type assertion can return two values; the first value (like `t` above, which is the value of the variable we are checking); and the second, which is a boolean indicator of whether the assertion succeeded.
```go
// type-lessons.go

...
var i interface{} = "a string"
t, ok := i.(string)
```

Our second variable, `ok`, is a boolean value that holds whether our type assertion was correct or not. So, in our example above, `ok` would be true because i is of type string.

However, this method is only useful for validating an intuition of a variable’s type without a panic being thrown. If we are unsure of the type of a variable, we can use our second option, which is called `Type switching`.

This is similar to a normal switch statement, which switches through possible types of a variable, rather than the normal switching of values. Here, we will extract the type from our interface variable and switch through several type options.
```go
// type-lessons.go

package main

func main() {
  var testVar interface{} = 42.56

  switch t := testVar.(type) {
    case string: 
      fmt.Println("Variable is of type string!")
    case int:
      fmt.Println("Variable is of type int!")
    case float32:
      fmt.Println("Variable is of type float32!")
    default:
      fmt.Println("Variable type unknown")
  }
}
```

# Type conversion in Go

Type conversion is simply changing a value from one type to another, but in Go there is a caveat, which is that types can only be converted between related or similar types. Let’s look at an example:
```go
// type-conversion.go

package main

import (
  "fmt"
  "strings"
)

type myString string

func (m myString) capitalize() myString {
    capStr := strings.ToUpper(string(m))

    return myString(capStr)
}

func main() {
    fmt.Println("Hello World!")

    var m myString = "test"

    fmt.Println(m.capitalize())
}
```

In our example above, we have defined a type `myString` using the basic type `string`. This means that `myString` inherits the data structure of `string`, but we can give it its own methods that type `string` will not have, like the method `capitalize` in our example.

We can also see in our example that we were able to convert our type `myString` to `string` and convert `string` to `myString`. This is because they share similar data structures.

Another example of types that can be converted explicitly (without using any special tricks or libraries) is `int` to `float64` and vice versa. Here’s a simple example:
```go
// type-conversion-int.go

package main

import "fmt"

func main() {
  var simpleInt int = 3

  var simpleFloat float64 = 4.5

  // fmt.Println("This will throw an error: ", simpleInt + simpleFloat)

  fmt.Println("This will work correctly: ", simpleInt + int(simpleFloat))

  fmt.Println("This will work correctly too: ", float64(simpleInt) + simpleFloat)
}
```

In our example above, we were able to convert between `int` and `float64` explicitly, just as we could between `myString` and `string`, because they both have similar data structures — but we cannot convert `string` to `int`.