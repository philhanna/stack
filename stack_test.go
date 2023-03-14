package stack

import (
	"encoding/xml"
	"flag"
	"fmt"
	"math"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStack is a parameterized set of test cases for stacks of various types:
//   - strings
//   - integers
//	 - float64
//	 - runes
//	 - bool
//   - structs
//   - functions
//	 - channels
//   - stacks
func TestStack(t *testing.T) {

	tests := []struct {
		name  string
		input []any
		want  []any
	}{
		// Test stacks of strings
		{"strings", []any{"Larry", "Curly", "Moe"}, []any{"Moe", "Curly", "Larry"}},

		// Test stacks of integers
		{"integers", []any{3, 1, 4, 2, 6}, []any{6, 2, 4, 1, 3}},

		// Test stacks of float64
		{"float64", []any{math.Pi, math.Phi, math.E, -17.3}, []any{-17.3, math.E, math.Phi, math.Pi}},

		// Test stacks of runes
		{"float64", []any{'a', 'b', 'c', '\u2023'}, []any{'\u2023', 'c', 'b', 'a'}},

		// Test stacks of bool
		{"bool", []any{true, false, false}, []any{false, false, true}},

		// Test stacks of structs
		{"tuples", []any{NewPoint(3, 4), NewPoint(-1, 17)}, []any{NewPoint(-1, 17), NewPoint(3, 4)}},

		// Test stacks of maps
		{"maps", []any{NewPointMap("5"), NewPointMap("3", "1", "2"), NewPointMap()}, []any{NewPointMap(), NewPointMap("3", "1", "2"), NewPointMap("5")}},

		// Test stacks of functions
		{"functions", []any{
			http.NotFound,
			xml.Escape,
			flag.Bool,
			math.Acos,
		}, []any{
			math.Acos,
			flag.Bool,
			xml.Escape,
			http.NotFound,
		},
		},

		// Test stacks of channels
		{"channels", []any{new(chan int), new(chan bool)}, []any{new(chan bool), new(chan int)}},
		
		// Test stacks of stacks
		{"stacks", []any{
			NewStack[rune](),
			NewStack[int](),
			NewStack[string](),
			NewStack[float64](),
		}, []any{
			NewStack[float64](),
			NewStack[string](),
			NewStack[int](),
			NewStack[rune](),
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[any]()
			wants := make([]any, 0)

			// Load the stack with the Push() method
			n := 0
			for i, item := range tt.input {
				stack.Push(item)
				wantItem := tt.want[i]
				wants = append(wants, wantItem)
				n++
			}

			// Test the Peek() method
			if n > 0 {
				want := wants[0]
				wantType := fmt.Sprintf("%T", want)
				have := stack.Peek()
				if strings.HasPrefix(wantType, "func") {
					// Cannot compare functions for equal
				} else {
					assert.Equal(t, want, have, fmt.Sprintf("Peek: want=%v,have=%v\n", want, have))
				}
			}

			// Test the Pop() method
			assert.Equal(t, n, stack.Len())
			for i := 0; i < n; i++ {
				want := wants[i]
				wantType := fmt.Sprintf("%T", want)
				have := stack.Pop()
				if strings.HasPrefix(wantType, "func") {
					// Cannot compare functions for equal
				} else {
					assert.Equal(t, want, have, fmt.Sprintf("Pop: want=%v,have=%v\n", want, have))
				}
			}

			// Stack should be empty now
			assert.True(t, stack.IsEmpty())

			// Test the Peek() and Pop() methods on an empty stack
			var empty any
			assert.Equal(t, empty, stack.Peek())
			assert.Equal(t, empty, stack.Pop())
		})
	}
}

// For testing a stack of structures
type Point struct {
	row, col int
}

func NewPoint(r, c int) Point {
	p := new(Point)
	p.row = r
	p.col = c
	return *p
}

// For testing a stack of maps
type PointMap map[Point]string

func NewPointMap(s ...string) PointMap {
	pm := make(map[Point]string)
	for i := 0; i < len(s); i++ {
		n := int(s[i][0])
		point := NewPoint(n, n+1)
		pm[point] = s[i]
	}
	return pm
}

