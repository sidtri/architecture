package main

// Stack
type Stack struct {
  items []*Node
}

func (s *Stack) Push(data *Node) {
	s.items = append(s.items, data)
}

func (s *Stack) IsEmpty() bool{
	return len(s.items) == 0
}

func (s *Stack) Pop() ([]*Node, *Node) {
	// if s.IsEmpty() {
	//	return nil, nil
	// }

	popElement := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items)-1]

	return s.items, popElement
}
