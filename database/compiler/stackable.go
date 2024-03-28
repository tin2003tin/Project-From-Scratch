package compiler

type SState struct {
	state int
}

func (s *SState) Value() interface{} {
	return s.state
}

type Nontermial struct {
	list [][]string
}

func (nt *Nontermial) Value() interface{} {
	return nt.list
}

type Terminal struct {
	id interface{}
}

func (t *Terminal) Value() interface{} {
	return t.id
}

func (s *State) Nontermial() interface{} {
	return s.Index
}
