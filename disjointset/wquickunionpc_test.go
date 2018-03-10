package goalgorithms

import (
	"reflect"
	"testing"
)

func TestNewWQuickUnionSetPC(t *testing.T) {
	wantSize := 15

	set := NewWQuickUnionSetPC(wantSize)
	if set == nil {
		t.Fatalf("NewWQuickUnionSetPC(%d) = nil, want something != nil", wantSize)
	}

	if got := set.Size(); got != wantSize {
		t.Errorf("NewWQuickUnionSetPC(%d) = %d, want %d", wantSize, got, wantSize)
	}
}

func TestWQuickUnionSetPC_Clone(t *testing.T) {
	want := WQuickUnionSetPC{[]int{0, 0, 1, 3, 3}, []int{3, 1, 1, 2, 1}}
	got := want.Clone()
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("%v.Clone() = %v, want %v", want, got, want)
	}
}

func TestWQuickUnionSetPC_Union(t *testing.T) {
	tests := []struct {
		name string
		set  WQuickUnionSetPC
		a, b int
		want WQuickUnionSetPC
	}{
		{"Connect 0 and 1", WQuickUnionSetPC{[]int{0, 1, 2, 3, 4}, []int{1, 1, 1, 1, 1}}, 0, 1, WQuickUnionSetPC{[]int{1, 1, 2, 3, 4}, []int{1, 2, 1, 1, 1}}},
		{"Connect smaller to larger", WQuickUnionSetPC{[]int{0, 0, 1, 3, 3}, []int{3, 1, 1, 2, 1}}, 2, 4, WQuickUnionSetPC{[]int{0, 0, 0, 0, 3}, []int{5, 1, 1, 2, 1}}},
		{"Connect nested", WQuickUnionSetPC{[]int{0, 0, 1, 2, 3}, []int{5, 1, 1, 1, 1}}, 3, 4, WQuickUnionSetPC{[]int{0, 0, 1, 1, 1}, []int{5, 1, 1, 1, 1}}},
		{"Connect already flat", WQuickUnionSetPC{[]int{0, 0, 0, 0, 0}, []int{5, 1, 1, 1, 1}}, 3, 4, WQuickUnionSetPC{[]int{0, 0, 0, 0, 0}, []int{5, 1, 1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Clone()
			got.Union(tt.a, tt.b)
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("WQuickUnionSetPC.Union(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestWQuickUnionSetPC_IsConnected(t *testing.T) {
	tests := []struct {
		name    string
		set     WQuickUnionSetPC
		a, b    int
		wantIs  bool
		wantIds []int
	}{
		{"two sublings are connected", WQuickUnionSetPC{[]int{3, 3, 2, 2, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 0, 1, true, []int{2, 2, 2, 2, 1, 5}},
		{"left is connected to root", WQuickUnionSetPC{[]int{3, 3, 0, 3, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 0, 3, true, []int{3, 3, 0, 3, 1, 5}},
		{"two branches are connected", WQuickUnionSetPC{[]int{1, 3, 0, 3, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 2, 4, true, []int{1, 3, 1, 3, 3, 5}},
		{"is connected to self", WQuickUnionSetPC{[]int{3, 3, 0, 3, 1, 5}, []int{2, 2, 1, 5, 1, 1}}, 5, 5, true, []int{3, 3, 0, 3, 1, 5}},
		{"disjoint components are not connected", WQuickUnionSetPC{[]int{3, 3, 0, 3, 1, 5, 5}, []int{2, 2, 1, 5, 1, 2, 1}}, 0, 6, false, []int{3, 3, 0, 3, 1, 5, 5}},
		{"disjoint element is not connected", WQuickUnionSetPC{[]int{3, 3, 0, 3, 1, 5, 5, 7}, []int{2, 2, 1, 5, 1, 2, 1, 1}}, 6, 7, false, []int{3, 3, 0, 3, 1, 5, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.IsConnected(tt.a, tt.b)
			if got != tt.wantIs {
				t.Errorf("WQuickUnionSetPC.IsConnected(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.wantIs)
			}

			if !reflect.DeepEqual(tt.set.ids, tt.wantIds) {
				t.Errorf("After IsConnected(%v, %v, %v), ids = %v, want %v", tt.set, tt.a, tt.b, tt.set.ids, tt.wantIds)
			}
		})
	}
}
