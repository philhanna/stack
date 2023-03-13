package stack

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Stack is a LIFO list for elements of type T
type Stack[T any] struct {
	list []T
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewStack creates a new stack of type T and returns a pointer to it.
func NewStack[T any]() *Stack[T] {
	p := new(Stack[T])
	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Len returns the number of items in the stack.
func (p *Stack[T]) Len() int {
	return len(p.list)
}

// IsEmpty returns true if the stack depth is zero.
func (p *Stack[T]) IsEmpty() bool {
	return p.Len() == 0
}

// Push appends an item of type T to the stack.
func (p *Stack[T]) Push(item T) {
	p.list = append(p.list, item)
}

// Pop removes the topmost item from the stack and returns it.  If the
// stack is empty, returns the zero form of type T.
func (p *Stack[T]) Pop() T {
	if p.Len() > 0 {
		result := p.list[p.Len()-1]
		p.list = p.list[:p.Len()-1]
		return result
	}
	return *(new(T))
}

// Peek returns the topmost item from the stack and returns it.  The
// stack is not modified.  If the stack is empty, returns the zero form
// of type T.
func (p *Stack[T]) Peek() T {
	if p.Len() > 0 {
		result := p.list[p.Len()-1]
		return result
	}
	return *(new(T))
}
