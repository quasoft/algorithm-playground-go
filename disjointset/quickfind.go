package goalgorithms

// Go implementation of quick find as explained by Sedgwick at
// https://www.coursera.org/learn/algorithms-part1/lecture/EcF3P/quick-find

// QuickFindSet represent a union find data structure.
// Indices are the number of the element.
// Values are the id of the components.
type QuickFindSet struct {
	ids []int
}

// NewQuickFindSet creates a new set with the specified size,
// with IDs equal to the number of the element.
func NewQuickFindSet(size int) *QuickFindSet {
	s := &QuickFindSet{make([]int, size, size)}
	for i := 0; i < s.Size(); i++ {
		s.SetID(i, i)
	}
	return s
}

// Size returns the number of elements in the set.
func (s *QuickFindSet) Size() int {
	return len(s.ids)
}

// ID returns the ID of the component to which a is connected.
func (s *QuickFindSet) ID(element int) int {
	return s.ids[element]
}

// SetID changes the ID of the specified element.
func (s *QuickFindSet) SetID(element, id int) {
	s.ids[element] = id
}

// Union creates a connection between the specified elements,
// by updating the ID of all elements equal to a, to match the
// ID of the b element.
func (s *QuickFindSet) Union(a, b int) {
	from := s.ID(a)
	to := s.ID(b)
	for i := 0; i < s.Size(); i++ {
		if s.ID(i) == from {
			s.SetID(i, to)
		}
	}
}

// IsConnected returns true if there is a connection between the
// specified elements.
func (s *QuickFindSet) IsConnected(a, b int) bool {
	return s.ID(a) == s.ID(b)
}
