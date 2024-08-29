package main

import (
	"fmt"
	"memocache/cache"
	"time"
)

func testLongLivedCache() {
	// Create a new Cache instance
	myCache := cache.New[string, int]()

	// Set key-value pairs in the cache
	myCache.Set("one", 1)
	myCache.Set("two", 2)
	myCache.Set("three", 3)

	// Retrieve values from the cache
	value, found := myCache.Get("two")
	if found {
		fmt.Printf("Value for key 'two': %v\n", value)
	} else {
		fmt.Println("Key 'two' not found in the cache")
	}

	// Pop a key from the cache
	poppedValue, found := myCache.Pop("three")
	if found {
		fmt.Printf("Popped value for key 'three': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'three' not found in the cache")
	}

	// Remove a key from the cache
	myCache.Remove("one")

	// Try to retrieve a removed key
	removedValue, found := myCache.Get("one")
	if found {
		fmt.Printf("Value for key 'one': %v\n", removedValue)
	} else {
		fmt.Println("Key 'one' not found in the cache (after removal)")
	}
}

func testExpiringCache() {
	// Create a new TTLCache instance
	myTTLCache := cache.NewTTL[string, int]()

	// Set key-value pairs with TTL in the cache
	myTTLCache.Set("one", 1, 5*time.Second)
	myTTLCache.Set("two", 2, 10*time.Second)
	myTTLCache.Set("three", 3, 15*time.Second)

	// Retrieve values from the cache
	value, found := myTTLCache.Get("two")
	if found {
		fmt.Printf("Value for key 'two': %v\n", value)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Wait for a while to allow some items to expire
	time.Sleep(7 * time.Second)

	// Try to retrieve an expired key
	expiredValue, found := myTTLCache.Get("one")
	if found {
		fmt.Printf("Value for key 'one': %v\n", expiredValue)
	} else {
		fmt.Println("Key 'one' not found in the cache or has expired")
	}

	// Pop a key from the cache
	poppedValue, found := myTTLCache.Pop("two")
	if found {
		fmt.Printf("Popped value for key 'two': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Remove a key from the cache
	myTTLCache.Remove("three")
}

func main() {
	//testLongLivedCache()
	testExpiringCache()
}
