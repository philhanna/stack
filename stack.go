package stack

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Stack[T any] struct {
	list []T
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

func NewStack[T any]() *Stack[T] {
	p := new(Stack[T])
	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

func (p *Stack[T]) Len() int {
	return len(p.list)
}

func (p *Stack[T]) Push(item T) {
	p.list = append(p.list, item)
}

func (p *Stack[T]) Pop() T {
	if p.Len() > 0 {
		result := p.list[p.Len()-1]
		p.list = p.list[:p.Len()-1]
		return result
	}
	return *(new(T))
}

func (p *Stack[T]) Peek() T {
	if p.Len() > 0 {
		result := p.list[p.Len()-1]
		return result
	}
	return *(new(T))
}