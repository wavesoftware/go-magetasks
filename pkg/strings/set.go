package strings

// NewSet creates a new Set from given strings.
func NewSet(ins ...string) Set {
	s := Set{}
	s.All(ins)
	return s
}

// Set is basic implementation of string based set.
type Set struct {
	set map[string]bool
}

// Add adds a string to the Set.
func (s *Set) Add(in string) {
	if s.set == nil {
		s.set = make(map[string]bool)
	}
	s.set[in] = true
}

// All adds all elements of a slice to the Set.
func (s *Set) All(ins []string) {
	for _, in := range ins {
		s.Add(in)
	}
}

// Equal checks if given Set is equal to the other given Set.
func (s Set) Equal(other Set) bool {
	if len(s.set) != len(other.set) {
		return false
	}
	for elem := range s.set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Contains checks if element exists within the Set.
func (s Set) Contains(elem string) bool {
	_, ok := s.set[elem]
	return ok
}
