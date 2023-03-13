package stack

import "testing"
import "github.com/stretchr/testify/assert"

func TestStack_PushPop(t *testing.T) {
	stack := NewStack()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")

	stooge := stack.Pop()
	assert.Equal(t, "Moe", stooge)
	stooge = stack.Pop()
	assert.Equal(t, "Curly", stooge)
	stooge = stack.Pop()
	assert.Equal(t, "Larry", stooge)

	assert.Equal(t, 0, stack.Len())
}
