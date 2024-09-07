package main

import (
	"testing"

	mock_main "github.com/favtuts/gomock-testing/mocks"
	"github.com/golang/mock/gomock"
)

func TestApiClientProcess(t *testing.T) {
	ctrl := gomock.NewController(t)

	// 👇 create new mock client
	mockApiClient := mock_main.NewMockApiClient(ctrl)

	// 👇 configure our mock `GetData` function to return mock data
	mockApiClient.EXPECT().GetData().Return("Hello World")

	length := Process(mockApiClient)

	if length != 11 {
		t.Fatalf("want: %d, got: %d\n", 11, length)
	}
}
