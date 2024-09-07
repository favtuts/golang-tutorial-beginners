package main

// example Fetcher interface that retrieves data from external API.
type Fetcher interface {
	FetchData() ([]byte, error)
}

// example function using the Fetcher interface
func ProcessFetcherData(client Fetcher) ([]byte, error) {
	data, err := client.FetchData()
	return data, err
}
