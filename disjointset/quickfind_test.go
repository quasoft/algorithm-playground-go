package goalgorithms

import (
	"reflect"
	"testing"
)

func TestNewDisjointSet(t *testing.T) {
	wantSize := 15

	set := NewDisjointSet(wantSize)
	if set == nil {
		t.Fatalf("NewDisjointSet(%d) = nil, want something != nil", wantSize)
	}

	if got := set.Size(); got != wantSize {
		t.Errorf("NewDisjointSet(%d) = %d, want %d", wantSize, got, wantSize)
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name string
		set  DisjointSet
		a, b int
		want DisjointSet
	}{
		{"Connect 0 and 1", DisjointSet{0, 1, 2, 3, 4, 5, 6, 7}, 0, 1, DisjointSet{1, 1, 2, 3, 4, 5, 6, 7}},
		{"Connect 1 and 0", DisjointSet{0, 1, 2, 3, 4, 5, 6, 7}, 1, 0, DisjointSet{0, 0, 2, 3, 4, 5, 6, 7}},
		{"Connect 3 and 4", DisjointSet{0, 1, 2, 3, 4, 5, 6, 7}, 3, 4, DisjointSet{0, 1, 2, 4, 4, 5, 6, 7}},
		{"Connect 0 and 7", DisjointSet{0, 1, 2, 3, 4, 5, 6, 7}, 0, 7, DisjointSet{7, 1, 2, 3, 4, 5, 6, 7}},
		{"Connect already connected", DisjointSet{0, 1, 2, 2, 4, 5, 6, 7}, 2, 3, DisjointSet{0, 1, 2, 2, 4, 5, 6, 7}},
		{"Connect two components", DisjointSet{0, 0, 3, 0, 4, 4, 5, 3}, 3, 2, DisjointSet{3, 3, 3, 3, 4, 4, 5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := make(DisjointSet, len(tt.set), len(tt.set))
			copy(set, tt.set)
			set.Union(tt.a, tt.b)
			got := set
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DisjointSet.Union(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsConnected(t *testing.T) {
	tests := []struct {
		name string
		set  DisjointSet
		a, b int
		want bool
	}{
		{"0 and 1 connected", DisjointSet{3, 3, 0, 5, 0, 5, 5, 5}, 0, 1, true},
		{"2 and 4 connected", DisjointSet{3, 3, 0, 5, 0, 5, 5, 5}, 2, 4, true},
		{"6 and 7 connected", DisjointSet{3, 3, 0, 5, 0, 5, 5, 5}, 6, 7, true},
		{"0 and 2 not connected", DisjointSet{3, 3, 0, 5, 0, 5, 5, 5}, 0, 2, false},
		{"2 and 7 not connected", DisjointSet{3, 3, 0, 5, 0, 5, 5, 5}, 2, 7, false},
		{"0 and 7 not connected", DisjointSet{3, 3, 0, 5, 0, 5, 5, 5}, 0, 7, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.IsConnected(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("DisjointSet.IsConnected(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}
