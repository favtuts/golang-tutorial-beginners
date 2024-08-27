# How GOROOT and GOPATH Works
* https://tuts.heomi.net/how-goroot-and-gopath-works/

# Ensure to use Go version 1.18.3

Throughout this entire article, I will be using the go version `1.18.3`:
```sh
$ goenv versions
$ goenv install -l
$ goenv install 1.18.3
$ goenv global 1.18.3
$ echo $GOROOT
$ echo $GOPATH
```

# Switch to Go Modules

We are going to turn on `GO111MODULE` from environment variables.
```sh
$ go env -w GO111MODULE=on
$ go env GO111MODULE
on
```

# Create Go workspace

First, let's start off with the `operations` package
```sh
DIR="$HOME/go-projects/golang-tutorial-beginners/go-modules-demo"   # My root workspace
mkdir -p $DIR/operations && cd $DIR/operations                      # Create package directory.
go mod init operations                                              # Initialize go module.
touch $DIR/operations/operations.go                                 # Create the source file.
```

Paste the following source in `operations.go`.
```go
package operations

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}
```

Then, let's write the driver application `calcapp`.
```sh
mkdir -p $DIR/calcapp $DIR/calcapp/formatters                     # Create package directory.
cd $DIR/calcapp && go mod init calcapp                            # Initialize go module.
touch $DIR/calcapp/formatters/formatters.go $DIR/calcapp/main.go  # Create source files.
```

Also, for the driver application, we need a third-party package called `chalk`.
```sh
go get github.com/ttacon/chalk      # Installs the chalk package.
go mod tidy                         # Sync sums
```

Paste the following source in `formatters.go`.
```go
package formatters

import (
	"fmt"
	"github.com/ttacon/chalk"
)

func Red(message string) {
	fmt.Println(
		chalk.Red,
		message,
		chalk.ResetColor,
	)

}

func Green(message string) {
	fmt.Println(
		chalk.Green,
		message,
		chalk.ResetColor,
	)
}
```

Now notice that we have a custom local package called `operations`. We need to import that package into our `calcapp` to make it work.

To point to the local version of a dependency in Go rather than the one over the web, we use the `replace` keyword within the `go.mod` file.

The `replace` line goes above your require statements, like so: 

**calcapp/go.mod**

```sh
module calcapp

go 1.18

// Can be github.com/yourname/operations
replace favtuts.com/operations => ../operations

// 💡 The actual semantic version hash will be different in yours.
require github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31

// This is a random version number I added. You can actually put any semantic version here
require yasint.dev/operations v0.0.0
```
