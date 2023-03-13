package stack

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Stack struct {
	list []string
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

func NewStack() *Stack {
	p := new(Stack)
	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

func (p *Stack) Len() int {
	return len(p.list)
}

func (p *Stack) Push(s string) {
	p.list = append(p.list, s)
}

func (p *Stack) Pop() string {
	if p.Len() > 0 {
		result := p.list[p.Len()-1]
		p.list = p.list[:p.Len()-1]
		return result
	}
	return ""
}

func (p *Stack) Peek() string {
	if p.Len() > 0 {
		result := p.list[p.Len()-1]
		return result
	}
	return ""
}