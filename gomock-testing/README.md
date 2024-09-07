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
$ go get -u github.com/stretchr/testify .
```

Install GoMock
```bash
$ go get github.com/golang/mock/gomock
```

Install MocGen CLI
Go version < 1.16
```bash
$ GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0
```

Go 1.16+
```bash
$ go install github.com/golang/mock/mockgen@v1.6.0
```

# Generate mock code

Let's say we have the following code in our `gomock_interface.go` file
```go
package main

type Fetcher interface {
	FetchData() ([]byte, error)
}

func ProcessFetcherData(client Fetcher) (int, error) {
	data, err := client.FetchData()
	if err != nil {
		return 0, err
	} else {
		return len(data), nil
	}
}
```

To generate mock implementation of ApiClient run the following code in the project root:
```bash
$ mockgen -source=gomock_interface.go -destination=mocks/gomock_interface.go
```

Here’s an example code snippet that illustrates how to create a mock object for our `Fetcher` interface using GoMock:
```go
package main

import (
	"bytes"
	"testing"

	mock_main "github.com/favtuts/gomock-testing/mocks"
	"github.com/golang/mock/gomock"
)

func TestFetchData(t *testing.T) {
	// Create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for the Fetcher interface
	mockFetcher := mock_main.NewMockFetcher(ctrl)

	// Set expectations on the mock object's behavior
	mockFetcher.EXPECT().FetchData().Return([]byte("data"), nil)

	// Call the code under test
	data, err := ProcessFetcherData(mockFetcher)

	// Assert the results
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !bytes.Equal(data, []byte("data")) {
		t.Errorf("Unexpected data: %v", data)
	}
}
```

Run the test:
```bash
$ go test -v -run=TestFetchData
=== RUN   TestFetchData
--- PASS: TestFetchData (0.00s)
PASS
ok      github.com/favtuts/gomock-testing       0.003s
```

