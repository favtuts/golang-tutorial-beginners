package main

import "fmt"

type Author struct {
	firstName, lastName, Biography string
	photoId                        int
}

type blogPostNestedStructFieldsExample struct {
	author    Author // nested struct field
	title     string
	postId    int
	published bool
}

type blogPostPromotedtructFieldsExample struct {
	Author    // directly passing the Author struct as a field - also called an anonymous field orembedded type
	title     string
	postId    int
	published bool
}

func demo_nested_struct_fields() {
	b := new(blogPostNestedStructFieldsExample)
	fmt.Println(b)

	b.author.firstName = "Alex"
	b.author.lastName = "Nnakwue"
	b.author.Biography = "I am a lazy engineer"
	b.author.photoId = 234333
	b.published = true
	b.title = "understand interface and struct type in Go"
	b.postId = 12345

	fmt.Println(*b)
}

func demo_promoted_struct_fields() {
	b := blogPostPromotedtructFieldsExample{
		Author:    Author{"Alex", "Nnakwue", "I am a lazy engineer", 234333},
		title:     "understand interface and struct type in Go",
		published: true,
		postId:    12345,
	}

	fmt.Println(b.firstName) // remember the firstName field is present on the Author struct?
	fmt.Println(b)

}
