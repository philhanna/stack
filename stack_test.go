package stack

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_string(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("A")
	stack.Push("B")
	assert.Equal(t, 2, stack.Len())

	var item string
	var err error

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		letter := "B"
		assert.Equal(t, item, letter)
	}

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		letter := "A"
		assert.Equal(t, item, letter)
	}
}

func TestStack_struct_contents_same(t *testing.T) {
	type Point struct {
		x, y int
	}
	stack := NewStack[Point]()
	stack.Push(Point{3, 4})
	if item, err := stack.Pop(); err != nil {
		t.Error(err)
	} else {
		wantX := 3
		wantY := 4
		assert.Equal(t, item.x, wantX)
		assert.Equal(t, item.y, wantY)
	}
}

func TestStack_map_contents_same(t *testing.T) {

	classic := make(map[string]int)
	classic["Larry"] = 2
	classic["Curly"] = 3
	classic["Moe"] = 1

	original := make(map[string]int)
	original["Moe"] = 1
	original["Larry"] = 2
	original["Shemp"] = 3

	noCurly := make(map[string]int)
	noCurly["Larry"] = 2
	noCurly["Moe"] = 1

	stack := NewStack[map[string]int]()
	stack.Push(original)
	stack.Push(noCurly)
	stack.Push(classic)

	assert.Equal(t, 3, stack.Len())

	var item map[string]int
	var err error

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		// Classic
		assert.Equal(t, 1, item["Moe"])
		assert.Equal(t, 2, item["Larry"])
		assert.Equal(t, 3, item["Curly"])
	}

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		// NoCurly
		assert.Equal(t, 1, item["Moe"])
		assert.Equal(t, 2, item["Larry"])
		_, ok := item["Curly"]
		assert.False(t, ok)
	}

	if item, err = stack.Pop(); err != nil {
		t.Error(err)
	} else {
		// Original
		assert.Equal(t, 1, item["Moe"])
		assert.Equal(t, 2, item["Larry"])
		assert.Equal(t, 3, item["Shemp"])
	}
}

func TestStack_Clear(t *testing.T) {
	stack := NewStack[float64]()
	stack.Push(math.Pi)
	stack.Push(math.E)
	assert.True(t, stack.Len() == 2)
	stack.Clear()
	assert.True(t, stack.IsEmpty())
}

func TestStack_FromSlice(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")
	list := stack.List
	assert.Equal(t, 3, len(list))
	assert.Equal(t, "Larry", list[0])
	assert.Equal(t, "Curly", list[1])
	assert.Equal(t, "Moe", list[2])
}

func TestStack_ToSlice(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("Larry")
	stack.Push("Curly")
	stack.Push("Moe")
	list := []string{
		"Joe",
		"Smith",
	}
	stack.List = list
	assert.Equal(t, 2, stack.Len())
	entry, _ := stack.Pop()
	assert.Equal(t, "Smith", entry)
	entry, _ = stack.Pop()
	assert.Equal(t, "Joe", entry)

}
