#  Go generics: Past designs and present release features 
* https://tuts.heomi.net/go-generics-past-designs-and-present-release-features/

# Init a Go project

Locate the project directory
```bash
$ mkdir go-generics
$ cd go-generics
```

Then create the module
```bash
$ go mod init github.com/favtuts/go-generics
```

# Reverse a list

```go
package main

import "fmt"

func main() {
        fmt.Println(reverse([]int{1, 2, 3, 4, 5}))
}

// T is a type parameter that is used like normal type inside the function
// any is a constraint on type i.e T has to implement "any" interface
func reverse[T any](s []T) []T {
        l := len(s)
        r := make([]T, l)

        for i, ele := range s {
                r[l-i-1] = ele
        }
        return r
}
```

Run the code:
```bash
$ go run .
[5 4 3 2 1]
```