# Exploring Go mocking methods and the GoMock framework
* https://tuts.heomi.net/exploring-go-mocking-methods-and-the-gomock-framework/

# Init a Go project

Locate the project directory
```bash
$ mkdir gomock-testing
$ cd gomock-testing
```

Then create the module
```bash
$ go mod init github.com/favtuts/gomock-testing
```


# Install dependencies

Install Testify
```bash
$ go get github.com/stretchr/testify
```

To update Testify to the latest version, use 
```bash
$ go get -u github.com/stretchr/testify.
```

Install GoMock
```bash
$ go get github.com/golang/mock/gomock
```