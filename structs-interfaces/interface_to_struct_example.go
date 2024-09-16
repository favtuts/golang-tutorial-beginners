package main

import "fmt"

type blogPostInterfaceToStructDemo struct {
	Data   interface{}
	postId int
}

func NewBlogPostStruct() interface{} {
	return &blogPostInterfaceToStructDemo{postId: 1234, Data: "Alexander"}
}

func demo_interface_to_struct_convert() {
	blogPost := NewBlogPostStruct().(*blogPostInterfaceToStructDemo)
	fmt.Println(blogPost.Data)
}
