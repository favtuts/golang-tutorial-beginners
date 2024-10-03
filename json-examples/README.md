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

## Nested Objects

Now, consider the case when you have a property called Dimensions, that measures the Height and Length of the bird in question:
```json
{
  "species": "pigeon",
  "description": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
```

To add a nested `dimensions` object, lets create a `dimensions` struct :
```go
type Dimensions struct {
  Height int
  Width int
}
```

Now, the `Bird` struct will include a `Dimensions` field:
```go
type Bird struct {
  Species string
  Description string
  Dimensions Dimensions
}
```

We can unmarshal this data using the same method as before:
```go
birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Println(bird)
// {pigeon likes to perch on rocks {24 10}}
```

Run the code:
```bash
$ go run json_to_nested_objects.go 
{pigeon likes to perch on rocks {24 10}}
```

## Primitive Types

We mostly deal with complex objects or arrays when working with JSON, but data like `3`, `3.1412` and `"birds"` are also valid JSON strings.

We can unmarshal these values to their corresponding data type in Go by using primitive types:
```go
numberJson := "3"
floatJson := "3.1412"
stringJson := `"bird"`

var n int
var pi float64
var str string

json.Unmarshal([]byte(numberJson), &n)
fmt.Println(n)
// 3

json.Unmarshal([]byte(floatJson), &pi)
fmt.Println(pi)
// 3.1412

json.Unmarshal([]byte(stringJson), &str)
fmt.Println(str)
// bird
```

Run the code:
```bash
$ go run json_to_primitive_types.go 
3
3.1412
bird
```

## Time Values

Did you know that if you try to decode an ISO 8601 date string like `2021-10-18T11:08:47.577Z` into a `time.Time` struct, it will work out of the box?
```go
dateJson := `"2021-10-18T11:08:47.577Z"`
var date time.Time
json.Unmarshal([]byte(dateJson), &date)

fmt.Println(date)
// 2021-10-18 11:08:47.577 +0000 UTC
```

Run the code:
```bash
$ go run json_to_time_values.go
2021-10-18 11:08:47.577 +0000 UTC
```

Here, `dateJson` is a JSON string type, but when we unmarshal it into a `time.Time` variable, it is able to understand the JSON data on its own. Well, this is because the `time.Time` struct has a custom [UnmarshalJSON](https://pkg.go.dev/time#Time.UnmarshalJSON) method that handles this case.

This will even work if the `time.Time` type is embedded within another struct:
```go
type Bird struct {
	Species     string
	Description string
	CreatedAt   time.Time
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "createdAt": "2021-10-18T11:08:47.577Z"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks 2021-10-18 11:08:47.577 +0000 UTC}
}
```

Run the code:
```bash
$ go run json_to_time_values.go 
{pigeon likes to perch on rocks 2021-10-18 11:08:47.577 +0000 UTC}
```

## Custom Parsing Logic

Similar to the `time.Time` struct, we can also create custom types that implement the [Unmarshaler](https://pkg.go.dev/encoding/json#Unmarshaler) interface. This will allow us to define custom logic for decoding JSON data into our custom types.

Suppose we receive the `dimensions` data as a formatted string:
```json
{
  "species": "pigeon",
  "description": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
```

We can modify the `Dimensions` type to implement the `Unmarshaler` interface, which will have custom parsing logic for our data:
```go
type Dimensions struct {
	Height int
	Width  int
}

// unmarshals a JSON string with format
// "heightxwidth" into a Dimensions struct
func (d *Dimensions) UnmarshalJSON(data []byte) error {
	// the "data" parameter is expected to be JSON string as a byte slice
	// for example, `"20x30"`

	if len(data) < 2 {
		return fmt.Errorf("dimensions string too short")
	}
	// remove the quotes
	s := string(data)[1 : len(data)-1]
	// split the string into its two parts
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return fmt.Errorf("dimensions string must contain two parts")
	}
	// convert the two parts into ints
	height, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("dimension height must be an int")
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("dimension width must be an int")
	}
	// assign the two ints to the Dimensions struct
	d.Height = height
	d.Width = width
	return nil
}
```

Now, if we try to unmarshal the JSON data, it will create a `Dimensions` struct with the correct values:
```go
birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "dimensions":"20x30"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Printf("%+v", bird)
// {Species:pigeon Description:likes to perch on rocks Dimensions:{Height:20 Width:30}}
```

Run the code:
```bash
$ go run json_implement_unmarshaler.go 
{Species:pigeon Description:likes to perch on rocks Dimensions:{Height:20 Width:30}}
```

# JSON Struct Tags – Custom Field Names

We saw earlier that Go uses convention to determine the attribute name for mapping JSON properties. Although sometimes, we want a different attribute name than the one provided in your JSON data. For example, consider the below data:
```json
{
  "birdType": "pigeon",
  "what it does": "likes to perch on rocks"
}
```

Here, we would prefer `birdType` to remain as the `Species` attribute in our Go code. It is also not possible for us to provide a suitable attribute name for a key like "what it does".

To solve this, we can use struct field tags:
```go
type Bird struct {
  Species string `json:"birdType"`
  Description string `json:"what it does"`
}
```

Now, we can explicitly tell our code which JSON property to map to which attribute.
```go
birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Println(bird)
// {pigeon likes to perch on rocks}
```

Run the code:
```bash
$ go run json_custom_field_names.go 
{pigeon likes to perch on rocks}
```