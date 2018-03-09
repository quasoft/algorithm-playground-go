package goalgorithms

import (
	"reflect"
	"testing"
)

func TestNewWQuickUnionSet(t *testing.T) {
	wantSize := 15

	set := NewWQuickUnionSet(wantSize)
	if set == nil {
		t.Fatalf("NewWQuickUnionSet(%d) = nil, want something != nil", wantSize)
	}

	if got := set.Size(); got != wantSize {
		t.Errorf("NewWQuickUnionSet(%d) = %d, want %d", wantSize, got, wantSize)
	}
}

func TestWQuickUnionSet_Clone(t *testing.T) {
	want := WQuickUnionSet{[]int{0, 0, 1, 3, 3}, []int{3, 1, 1, 2, 1}}
	got := want.Clone()
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("%v.Clone() = %v, want %v", want, got, want)
	}
}

func TestWQuickUnionSet_Union(t *testing.T) {
	tests := []struct {
		name string
		set  WQuickUnionSet
		a, b int
		want WQuickUnionSet
	}{
		{"Connect 0 and 1", WQuickUnionSet{[]int{0, 1, 2, 3, 4}, []int{1, 1, 1, 1, 1}}, 0, 1, WQuickUnionSet{[]int{1, 1, 2, 3, 4}, []int{1, 2, 1, 1, 1}}},
		{"Connect smaller to larger", WQuickUnionSet{[]int{0, 0, 1, 3, 3}, []int{3, 1, 1, 2, 1}}, 1, 4, WQuickUnionSet{[]int{0, 0, 1, 0, 3}, []int{5, 1, 1, 2, 1}}},
		{"Connect already connected", WQuickUnionSet{[]int{0, 1, 2, 2, 4}, []int{1, 1, 2, 1, 1}}, 2, 3, WQuickUnionSet{[]int{0, 1, 2, 2, 4}, []int{1, 1, 2, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Clone()
			got.Union(tt.a, tt.b)
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("WQuickUnionSet.Union(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestWQuickUnionSet_IsConnected(t *testing.T) {
	tests := []struct {
		name string
		set  WQuickUnionSet
		a, b int
		want bool
	}{
		{"two sublings are connected", WQuickUnionSet{[]int{3, 3, 0, 3, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 0, 1, true},
		{"left is connected to root", WQuickUnionSet{[]int{3, 3, 0, 3, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 0, 3, true},
		{"two branches are connected", WQuickUnionSet{[]int{3, 3, 0, 3, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 2, 4, true},
		{"is connected to self", WQuickUnionSet{[]int{3, 3, 0, 3, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 5, 5, true},
		{"disjoint components are not connected", WQuickUnionSet{[]int{3, 3, 0, 3, 1, 5, 5}, []int{2, 2, 1, 5, 1, 2, 1}}, 0, 6, false},
		{"disjoint element is not connected", WQuickUnionSet{[]int{3, 3, 0, 3, 1, 5, 5, 7}, []int{2, 2, 1, 5, 1, 2, 1, 1}}, 6, 7, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.IsConnected(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("WQuickUnionSet.IsConnected(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}
