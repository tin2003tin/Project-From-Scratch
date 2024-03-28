package lib

type Stack struct {
	Item []Stackable
}

type Stackable interface {
	Value() interface{}
}

func (s *Stack) Push(ss Stackable) {
	s.Item = append(s.Item, ss)
}

func (s *Stack) Pop() Stackable {
	if len(s.Item) == 0 {
		return nil
	}
	top := s.Item[len(s.Item)-1]
	s.Item = s.Item[:len(s.Item)-1]
	return top
}

func (s *Stack) Top() Stackable {
	return s.Item[len(s.Item)-1]
}
