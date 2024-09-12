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
