package main

import "fmt"

// interface definition
type Publisher interface {
	Publish() error
}

// method with a value receiver
func (b blogPost) Publish() error {
	fmt.Printf("The title on %s has been published by %s, with postId %d\n", b.title, b.author, b.postId)
	return nil
}

// Receives any type that satifies the Publisher interface
func PublishPost(publish Publisher) error {
	return publish.Publish()
}

func demo_publisher_interafce_testing() {
	b := blogPost{"Alex", "understanding structs and interface types", 12345}

	fmt.Println(b.Publish())

	d := &b // pointer receiver for the struct type

	b.author = "Chinedu"

	fmt.Println(d.Publish())
}

func demo_interface_type_testing() {
	var p Publisher
	fmt.Println(p)

	p = blogPost{"Alex", "understanding structs and interface types", 12345}

	fmt.Println(p.Publish())
}

func demo_interface_type_function_argument() {
	var p Publisher

	fmt.Println(p)

	b := blogPost{"Alex", "understand structs and interface types", 12345}

	fmt.Println(b)

	PublishPost(b)
}
