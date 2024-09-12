package main

import "fmt"

type blogPost struct {
	author string // field
	title  string // field
	postId int    // field
}

func NewBlogPost() *blogPost {
	return &blogPost{
		author: "Alexander",
		title:  "Learning structs and interfaces in Go",
		postId: 4555,
	}
}

func demo_init_struct_use_literal() {
	var b blogPost // initialize the struct type
	fmt.Println(b) // print the zero value: {  0}

	newBlogPost := *NewBlogPost()
	fmt.Println(newBlogPost)

	// alternatively
	b = blogPost{
		author: "Alex",
		title:  "Understand struct and interface types",
		postId: 12345,
	}
	fmt.Println(b)
}

func demo_access_struct_fields_use_dot() {
	var b blogPost // b is a type Alias for the BlogPost
	b.author = "Alex"
	b.title = "understand structs and interface types"
	b.postId = 12345

	fmt.Println(b)

	b.author = "Chinedu" // since everything is pass by value by default in Go, we can update this field after initializing - see pointer types later

	fmt.Println("Updated Author's name is: ", b.author)
}
