package main

type ItemType interface{}

type Stack struct {
	items []ItemType
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item ItemType) {
	if s.items == nil {
		s.items = []ItemType{}
	}

	s.items = append(s.items, item)
}

func (s *Stack) Pop() *ItemType {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return &item
}
