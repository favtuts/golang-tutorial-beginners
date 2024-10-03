# Using Maps in Golang – With Examples 
* https://tuts.heomi.net/using-maps-in-golang-with-examples/

# Init a Go project

Locate the project directory
```bash
$ mkdir map-examples
$ cd map-examples
```

Then create the module
```bash
$ go mod init github.com/favtuts/map-examples
```

# Creating a Map

When we want to create a new map, we need to declare its key and value types. The syntax for declaring a map type is:
```go
var m map[keyType]valueType
```

This only declares the type, but to create an actual map, we need to use the built-in `make` function.

For example, if we want to create a map with `string` keys and `int` values, we can do so as follows:
```go
myMap := make(map[string]int)
```

This creates an empty map with `string` keys and `int` values.

# Adding Key-Value Pairs to a Map

We can add key-value pairs to a map using the following syntax:
```go
m[key] = value
```

This adds a key-value pair to the map `m`. If the key already exists in the map, then the value is updated. If the key does not exist, then a new key-value pair is added to the map.

For example, if we want to add the key-value pair `"the_answer": 42` to the map `myMap`, we can do so as follows:
```go
myMap["the_answer"] = 42
```

If you set a key that already exists, then it’s previous value is overwritten.

# Getting a Value from a Map

We can get the value associated with a key from a map using the following syntax:
```bash
value = m[key]
```

Let’s see an example of this. If we want to get the value associated with the key `"the_answer"` from the map `myMap`, we can do so as follows:
```go
value := myMap["the_answer"]
```

However, if we try to get the value associated with the key `"the_question"` from the map `myMap`, we will get the zero value of the value type, which is 0 in this case.
```go
value := myMap["the_question"]
// value == 0
```

# Checking if a Key Exists in a Map

Since we get a zero value when a key doesn’t exist in a map, we can’t use this to determine if the key was present in the map or not. Fortunately, when obtaining a value from a map, we can also get a boolean value that indicates whether the key was present in the map or not:
```go
value, ok := m[key]
```

The value of the second variable `ok` is `true` if the key was present in the map, and `false` otherwise.

# Deleting Key-Value Pairs

We can delete a key-value pair from a map using the inbuilt `delete` function:
```go
delete(m, key)
```

For example, if we want to delete the key-value pair `"the_answer": 42` from the map `myMap`, we can do so as follows:
```go
delete(myMap, "the_answer")
```

The `delete` function works even if the key doesn’t exist in the map, or even if the map is `nil`. In that case, it does nothing, and acts as a no-op.

# Iterating Over a Map

We can iterate over a map using the `range` keyword. In each iteration, we can access a the key and value variables:
```go
for key, value := range m {
    // do something with key and value
}
```

# Run the full code:
```bash
$ go run map_operations.go 
The value of key `the_answer` is 42
Can not find the value of key `the_new_answer`
the_answer has value: 42
```

# Performance of Various Map Operations

We have the [benchmark](https://www.practical-go-lessons.com/chap-34-benchmarks) codes and after running this on a laptop, i got the followings results:

| Operation	| Time Taken |
| ----------| ---------- |
| Map Initialization |	3.874 ns/op |
| Add Key-Value Pair (to an empty map) | 5.339 ns/op |
| Get Value	| 6.99 ns/op |
| Update Value | 10.486 ns/op |
| Delete Key-Value Pair | 12.71 ns/op |
| Add Key-Value Pair (to a large map) |	123 ns/op |
| Get Value (to a large map) | 7.6 ns/op |
| Update Value (to a large map) | 8.8 ns/op |
| Delete Key-Value Pair (to a large map) | 7.9 ns/op |


To run all benchmarks in a module, use the command : 
```bash
$ go test -bench=.
```

To run a specific benchmark, use this command : 
```bash
$ go test -bench ConcatenateBuffer
```

To display memory statistics, add the flag “benchmem” : 
```bash
$ go test -bench . -benchmem
```

Run the benchmark:
```bash
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/favtuts/map-examples
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkMapInit-8              180867490                6.778 ns/op
BenchmarkMapGet-8               73969592                15.66 ns/op
BenchmarkMapAdd-8               77217672                13.40 ns/op
BenchmarkMapUpdate-8            57868516                19.20 ns/op
BenchmarkMapDelete-8            44933188                25.43 ns/op
BenchmarkLargeMapGet-8          173768650                6.918 ns/op
BenchmarkLargeMapAdd-8           6417300               175.4 ns/op
BenchmarkLargeMapUpdate-8       69517159                18.07 ns/op
BenchmarkLargeMapDelete-8       100000000               10.63 ns/op
PASS
ok      github.com/favtuts/map-examples 12.834s
```