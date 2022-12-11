package jolt

type StringSet map[string]bool

func NewStringSet(refs []JobRef) *StringSet {
	s := &StringSet{}

	for _, ref := range refs {
		s.Add(string(ref))
	}

	return s
}

func (s *StringSet) Add(v string) {
	(*s)[v] = true
}

func (s *StringSet) Delete(v string) {
	delete((*s), v)
}

func (s *StringSet) Contains(v string) bool {
	_, ok := (*s)[v]
	return ok
}
