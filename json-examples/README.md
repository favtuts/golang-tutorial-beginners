#  A Complete Guide to JSON in Golang (With Examples) 
* https://tuts.heomi.net/a-complete-guide-to-json-in-golang-with-examples/

# Init a Go project

Locate the project directory
```bash
$ mkdir json-examples
$ cd json-examples
```

Then create the module
```bash
$ go mod init github.com/favtuts/json-examples
```


# Introduction

JSON is used as the de-facto standard for data serialization in many applications, ranging from REST APIs to configuration management. 

In this part, you’ll get familiar with how to marshal (encode) and unmarshal (decode) JSON in Go. We will learn how to convert from JSON raw data (strings or bytes) into Go types like structs, arrays, and slices, as well as unstructured data like maps and empty interfaces.

There are two types of data you will encounter when working with JSON:

* Structured data: structs, arrays, slices
* Unstructured data: maps, empty interfaces

# Decoding JSON Into Structs (Structured Data)

## Struct Object

“Structured data” refers to data where you know the format beforehand. For example, let’s say you have a bird object, where each bird has a `species` field and a `description` field :
```json
{
  "species": "pigeon",
  "description": "likes to perch on rocks"
}
```

To work with this kind of data, create a `struct` that mirrors the data you want to parse. In our case, we will create a bird struct which has a `Species` and `Description` attribute:
```go
type Bird struct {
  Species string
  Description string
}
```

And unmarshal it as follows:
```go
birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
var bird Bird	
json.Unmarshal([]byte(birdJson), &bird)
fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
//Species: pigeon, Description: likes to perch on rocks
```

By convention, Go uses the same title cased attribute names as are present in the case insensitive JSON properties. So the `Species` attribute in our `Bird` struct will map to the `species`, or `Species` or `sPeCiEs` JSON property.

Run the code:
```bash
$ go run json_to_struct_object.go
Species: pigeon, Description: likes to perch on rocks
```


## Array of objects

Let’s look at how we can decode an array of objects, like below:
```json
[
  {
    "species": "pigeon",
    "description": "likes to perch on rocks"
  },
  {
    "species":"eagle",
    "description":"bird of prey"
  }
]
```

Since each element of the array has the structure of the Bird struct, you can unmarshal it by creating a slice of birds :
```bash
birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
var birds []Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf("Birds : %+v", birds)
//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
```

Run the code:
```bash
$ go run json_to_array_objects.go 
Birds : [{Species:pigeon Description:likes to perch on rocks} {Species:eagle Description:bird of prey}]
```