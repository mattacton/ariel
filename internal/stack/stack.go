package stack

type Stack[T any] struct {
	Items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.Items) == 0 {
		panic("stack is empty")
	}
	i := len(s.Items) - 1
	item := s.Items[i]
	s.Items = s.Items[:i]
	return item
}

func (s *Stack[T]) Peek() T {
	if len(s.Items) == 0 {
		panic("stack is empty")
	}
	return s.Items[len(s.Items)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(s.Items) == 0
}
