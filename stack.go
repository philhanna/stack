package stack

import "errors"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Stack is a LIFO list of items.
type Stack[T any] struct {
	list []T
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewStack creates a new stack and returns it.
func NewStack[T any]() Stack[T] {
	pst := new(Stack[T])
	return *pst
}

// ---------------------------------------------------------------------
// Error messages
// ---------------------------------------------------------------------

var ErrorStackEmpty = errors.New("Stack empty")

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Len returns the number of items in the stack.
func (pst *Stack[T]) Len() int {
	return len(pst.list)
}

// IsEmpty returns true if the stack depth is zero.
func (pst *Stack[T]) IsEmpty() bool {
	return pst.Len() == 0
}

// Push appends an item onto the stack.
func (pst *Stack[T]) Push(item T) {
	pst.list = append(pst.list, item)
}

// Pop removes the topmost item from the stack and returns it. If the
// stack is empty, returns an error
func (pst *Stack[T]) Pop() (T, error) {
	empty := new(T)
	if pst.IsEmpty() {
		return *empty, ErrorStackEmpty
	}
	last := pst.Len() - 1
	result := pst.list[last]
	pst.list = pst.list[:last]
	return result, nil
}

// Peek returns the topmost item from the stack and returns it. The
// stack is not modified. If the stack is empty, returns an error.
func (pst *Stack[T]) Peek() (T, error) {
	empty := new(T)
	if pst.IsEmpty() {
		return *empty, ErrorStackEmpty
	}
	last := pst.Len() - 1
	result := pst.list[last]
	return result, nil
}

// Clear removes all elements from the stack
func (pst *Stack[T]) Clear() {
	for !pst.IsEmpty() {
		pst.Pop()
	}
}
