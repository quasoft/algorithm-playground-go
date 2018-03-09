package goalgorithms

import (
	"reflect"
	"testing"
)

func TestNewQuickUnionSet(t *testing.T) {
	wantSize := 15

	set := NewQuickUnionSet(wantSize)
	if set == nil {
		t.Fatalf("NewQuickUnionSet(%d) = nil, want something != nil", wantSize)
	}

	if got := set.Size(); got != wantSize {
		t.Errorf("NewQuickUnionSet(%d) = %d, want %d", wantSize, got, wantSize)
	}
}

func TestQuickUnionSet_Clone(t *testing.T) {
	want := QuickUnionSet{[]int{0, 1, 2, 3, 4}}
	got := want.Clone()
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("%v.Clone() = %v, want %v", want, got, want)
	}
}

func TestQuickUnionSet_Union(t *testing.T) {
	tests := []struct {
		name string
		set  QuickUnionSet
		a, b int
		want QuickUnionSet
	}{
		{"Connect 0 and 1", QuickUnionSet{[]int{0, 1, 2, 3, 4}}, 0, 1, QuickUnionSet{[]int{1, 1, 2, 3, 4}}},
		{"Connect 4 and 0", QuickUnionSet{[]int{0, 1, 2, 3, 4}}, 4, 0, QuickUnionSet{[]int{0, 1, 2, 3, 0}}},
		{"Connect already connected", QuickUnionSet{[]int{0, 1, 2, 2, 4}}, 3, 2, QuickUnionSet{[]int{0, 1, 2, 2, 4}}},
		{"Connect two components", QuickUnionSet{[]int{0, 0, 1, 3, 3}}, 2, 4, QuickUnionSet{[]int{3, 0, 1, 3, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Clone()
			got.Union(tt.a, tt.b)
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("QuickUnionSet.Union(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestQuickUnionSet_IsConnected(t *testing.T) {
	tests := []struct {
		name string
		set  QuickUnionSet
		a, b int
		want bool
	}{
		{"two sublings are connected", QuickUnionSet{[]int{3, 3, 0, 3, 1, 5}}, 0, 1, true},
		{"left is connected to root", QuickUnionSet{[]int{3, 3, 0, 3, 1, 5}}, 0, 3, true},
		{"two branches are connected", QuickUnionSet{[]int{3, 3, 0, 3, 1, 5}}, 2, 4, true},
		{"is connected to self", QuickUnionSet{[]int{3, 3, 0, 3, 1, 5}}, 5, 5, true},
		{"disjoint components are not connected", QuickUnionSet{[]int{3, 3, 0, 3, 1, 5, 5}}, 0, 6, false},
		{"disjoint element is not connected", QuickUnionSet{[]int{3, 3, 0, 3, 1, 5, 5, 7}}, 6, 7, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.IsConnected(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("QuickUnionSet.IsConnected(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}
