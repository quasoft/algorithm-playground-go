package goalgorithms

// Go implementation of quick find as explained by Sedgwick at
// https://www.coursera.org/learn/algorithms-part1/lecture/EcF3P/quick-find

// WQuickUnionSetPC represent a data structure for weighted union find with
// one pass path compression.
type WQuickUnionSetPC struct {
	// Indices are the numbers of the elements.
	// Values are the parent ids of the components.
	ids []int
	// Indices are the numbers of the elements.
	// Values are the weight of the component (number of elements inside the component).
	weight []int
}

// NewWQuickUnionSetPC creates a new set with the specified size,
// with IDs equal to the number of the element.
func NewWQuickUnionSetPC(size int) *WQuickUnionSetPC {
	s := &WQuickUnionSetPC{make([]int, size, size), make([]int, size, size)}
	for i := 0; i < s.Size(); i++ {
		s.SetID(i, i)
		s.SetWeight(i, 1)
	}
	return s
}

// Size returns the number of elements in the set.
func (s *WQuickUnionSetPC) Size() int {
	return len(s.ids)
}

// Clone creates a deep copy of WQuickUnionSetPC.
func (s *WQuickUnionSetPC) Clone() *WQuickUnionSetPC {
	new := NewWQuickUnionSetPC(s.Size())
	copy(new.ids, s.ids)
	copy(new.weight, s.weight)
	return new
}

// Root returns the root ID of the component to which the element is connected.
func (s *WQuickUnionSetPC) Root(element int) int {
	parent := element
	for s.ids[parent] != parent {
		s.ids[parent] = s.ids[s.ids[parent]]
		parent = s.ids[parent]
	}
	return parent
}

// SetID changes the ID of the specified element.
func (s *WQuickUnionSetPC) SetID(element, id int) {
	s.ids[element] = id
}

// Weight returns the number of elements inside the component rooted at the specified element.
func (s *WQuickUnionSetPC) Weight(element int) int {
	return s.weight[element]
}

// SetWeight changes the weight value for the specified element.
func (s *WQuickUnionSetPC) SetWeight(element, weight int) {
	s.weight[element] = weight
}

// Union creates a connection between the specified elements,
// by updating the ID of all elements equal to a, to match the
// ID of the b element.
func (s *WQuickUnionSetPC) Union(a, b int) {
	from := s.Root(a)
	to := s.Root(b)

	if from == to {
		// Already connected
		return
	}

	// Always connect smaller component to larger one
	if s.Weight(from) > s.Weight(to) {
		from, to = to, from
	}
	s.SetID(from, to)
	s.SetWeight(to, s.Weight(to)+s.Weight(from))
}

// IsConnected returns true if there is a connection between the
// specified elements.
func (s *WQuickUnionSetPC) IsConnected(a, b int) bool {
	return s.Root(a) == s.Root(b)
}
