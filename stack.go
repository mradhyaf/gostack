package gostack

type Stack[T any] []T

func (s *Stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		return *new(T), false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack[T]) Pop() (T, bool) {
	e, ok := s.Peek()
	if !ok {
		return e, false
	}
	*s = (*s)[:len(*s)-1]
	return e, true
}

func (s *Stack[T]) Push(e T) {
	*s = append(*s, e)
}
