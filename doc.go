/*
Stack is a library for using LIFO stacks of objects of any type.
It has all the familiar operations:
 Len(), which returns the number of items in the stack
 IsEmpty(), which returns true if the stack has no items
 Push(item any), which adds an item to the stack
 Pop(), which removes and returns an element from the stack
 Peek(), which returns the element on the top of the stack without removing it

Pop() and Peek() do not fail when called on an empty stack, they just
return the zero value of the stack item type. This is for convenience
in not returning two values.
*/
package stack