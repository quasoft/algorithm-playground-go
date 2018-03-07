package goalgorithms

// Go implementation of quick find as explained by Sedgwick at
// https://www.coursera.org/learn/algorithms-part1/lecture/EcF3P/quick-find

// QuickUnionSet represent a union find data structure.
// Indices are the numbers of the elements.
// Values are the ids of the components.
type QuickUnionSet struct {
	ids []int
}

// NewQuickUnionSet creates a new set with the specified size,
// with IDs equal to the number of the element.
func NewQuickUnionSet(size int) *QuickUnionSet {
	s := &QuickUnionSet{make([]int, size, size)}
	for i := 0; i < s.Size(); i++ {
		s.SetID(i, i)
	}
	return s
}

// Size returns the number of elements in the set.
func (s *QuickUnionSet) Size() int {
	return len(s.ids)
}

// Clone creates a deep copy of QuickUnionSet.
func (s *QuickUnionSet) Clone() *QuickUnionSet {
	new := NewQuickUnionSet(s.Size())
	copy(new.ids, s.ids)
	return new
}

// Root returns the root ID of the component to which the element is connected.
func (s *QuickUnionSet) Root(element int) int {
	parent := element
	for s.ids[parent] != parent {
		parent = s.ids[parent]
	}
	return parent
}

// SetID changes the ID of the specified element.
func (s *QuickUnionSet) SetID(element, id int) {
	s.ids[element] = id
}

// Union creates a connection between the specified elements,
// by updating the ID of all elements equal to a, to match the
// ID of the b element.
func (s *QuickUnionSet) Union(a, b int) {
	from := s.Root(a)
	to := s.Root(b)
	s.SetID(from, to)
}

// IsConnected returns true if there is a connection between the
// specified elements.
func (s *QuickUnionSet) IsConnected(a, b int) bool {
	return s.Root(a) == s.Root(b)
}
