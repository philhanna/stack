package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_String(t *testing.T) {
	var stooge string

	stack := NewStack[string]()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")

	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 3, stack.Len())
	assert.Equal(t, "Moe", stack.Peek())
	stooge = stack.Pop()
	assert.Equal(t, "Moe", stooge)
	stooge = stack.Pop()
	assert.Equal(t, "Curly", stooge)
	stooge = stack.Pop()
	assert.Equal(t, "Larry", stooge)
	assert.True(t, stack.IsEmpty())
	stooge = stack.Pop()
	assert.Equal(t, "", stooge)
	assert.True(t, stack.IsEmpty())
}

func TestStack_Int(t *testing.T) {

	var digit int

	stack := NewStack[int]()
	stack.Push(3)
	stack.Push(1)
	stack.Push(4)

	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 3, stack.Len())
	assert.Equal(t, 4, stack.Peek())
	digit = stack.Pop()
	assert.Equal(t, 4, digit)
	digit = stack.Pop()
	assert.Equal(t, 1, digit)
	digit = stack.Pop()
	assert.Equal(t, 3, digit)
	assert.True(t, stack.IsEmpty())
	digit = stack.Pop()
	assert.Equal(t, 0, digit)
	assert.True(t, stack.IsEmpty())
}
