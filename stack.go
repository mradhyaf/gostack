// Package gostack implements a generic stack.
package gostack

// A Stack[T] is a stack of elements of type T.
type Stack[T any] []T

// Peek returns the element from the top of the Stack and true if Stack is not empty.
// Otherwise, returns the zero-value of T and false.
func (s *Stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		return *new(T), false
	}
	return (*s)[len(*s)-1], true
}

// Pop removes, from the top of Stack, the element e.
// Returns e and true if Stack is not empty.
// Otherwise, returns the zero-value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	e, ok := s.Peek()
	if !ok {
		return e, false
	}
	*s = (*s)[:len(*s)-1]
	return e, true
}

// Push appends an element e of type T to the top of Stack.
func (s *Stack[T]) Push(e T) {
	*s = append(*s, e)
}
