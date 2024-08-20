package gostack

import (
	"fmt"
	"testing"
)

func TestPeek(t *testing.T) {
	s := Stack[int]{1}
	e, ok := s.Peek()
	if !ok || e != 1 {
		t.Fatalf("Peek() = %v, %v, expected 1", e, ok)
	} else if len(s) != 1 {
		t.Fatalf("len(IntStack) = %v, expected 1", len(s))
	}
}

func TestPop(t *testing.T) {
	s := Stack[int]{1}
	e, ok := s.Pop()
	if !ok || e != 1 {
		t.Fatalf("Pop() = %v, %v, expected 1", e, ok)
	} else if len(s) != 0 {
		t.Fatalf("len(IntStack) = %v, expected 0", len(s))
	}
}

func TestPush(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	if s[0] != 1 {
		t.Fatalf("len(IntStack) = %v, expected 1", len(s))
	}
}

func Example() {
	// Create and populate a new Stack[rune]
	s := Stack[rune]{}
	s.Push('a')
	s.Push('b')
	s.Push('d')

	// Pop all of the elements in s
	for e, ok := s.Pop(); ok; e, ok = s.Pop() {
		fmt.Println(string(e))
	}

	// Output:
	// d
	// b
	// a
}
