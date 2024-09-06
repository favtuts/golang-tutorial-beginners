package main

func ProcessData(data string, f func(string) string) string {
	result := f(data)
	// do some processing with result
	return result
}
