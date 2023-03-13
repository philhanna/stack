package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_String(t *testing.T) {

	stack := NewStack[string]()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")

	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 3, stack.Len())
	assert.Equal(t, "Moe", stack.Peek())
	assert.Equal(t, "Moe", stack.Pop())
	assert.Equal(t, "Curly", stack.Pop())
	assert.Equal(t, "Larry", stack.Pop())
	assert.True(t, stack.IsEmpty())
	assert.Equal(t, "", stack.Pop())
	assert.Equal(t, "", stack.Peek())
	assert.True(t, stack.IsEmpty())
}

func TestStack_Int(t *testing.T) {

	stack := NewStack[int]()
	stack.Push(3)
	stack.Push(1)
	stack.Push(4)

	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 3, stack.Len())
	assert.Equal(t, 4, stack.Peek())
	assert.Equal(t, 4, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, 3, stack.Pop())
	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.Pop())
	assert.Equal(t, 0, stack.Peek())
	assert.True(t, stack.IsEmpty())
}
