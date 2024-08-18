package gostack

import "testing"

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
