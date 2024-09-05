# Go long by generating PDFs in Golang with Maroto
* https://tuts.heomi.net/go-long-by-generating-pdfs-in-golang-with-maroto/

# Ensure latest Golang version

I am using [goenv](https://tuts.heomi.net/installing-multiple-version-of-golang-using-goenv/) to manage Go version in my machine.

```bash
# check all version which can be installed
$ goenv install -l

# install the latest version
$ goenv install 1.22.4

# make this version to be global
$ goenv global 1.22.4
$ goenv versions

# restart shell
$ exec $SHELL

# check go version
$ go version
go version go1.22.4 linux/amd64
```

# Initializing a new Go project

Locate the project location
```bash
$ mkdir pdf-gen-maroto
$ cd pdf-gen-maroto
```

Follow the following command:
```bash
$ go mod init github.com/USERNAME/PROJECT_NAME
```
In the command above, replace `github.com` with the domain where you store your files, such as `Bitbucket` or `Gitlab`. Additionally, substitute `USERNAME` with your username and `PROJECT_NAME` with the desired project name.

Run the command for init Go project
```bash
$ go mod init github.com/favtuts/maroto-pdf
```

To install Maroto as a dependency, run the following command:
```bash
$ go get github.com/johnfercher/maroto/v2@v2.1.2
```

For writing code, create a new file called main.go to the root of your project folder, and paste the following code into it:
```go
package main

import "fmt"

func main() {
 fmt.Println("Hello, World!")
}
```

Now, run the command `go run main.go` from the terminal, and check if it prints `Hello, World!` in the terminal. If it does, it means you are ready to get started.