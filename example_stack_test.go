package stack_test

import (
	"fmt"
	"github.com/philhanna/stack"
)

// ExampleStack illustrates how to use the Stack type.
func ExampleStack() {

	var item string

	// Create a stack for strings
	sst := stack.NewStack[string]()

	// Push three strings on the stack
	sst.Push("A")
	sst.Push("B")
	sst.Push("C")

	// Get the stack length
	length := sst.Len()
	fmt.Println(length)

	// Peek returns the top element without altering the stack
	item, _ = sst.Peek()
	fmt.Printf("%s,%d\n", item, sst.Len())

	// Pop returns the top element and removes it from the stack
	for !sst.IsEmpty() {
		item, _ = sst.Pop()
		fmt.Println(item)
	}

	// Output:
	// 3
	// C,3
	// C
	// B
	// A
}
