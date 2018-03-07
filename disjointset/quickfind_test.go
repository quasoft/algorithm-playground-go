package goalgorithms

import (
	"reflect"
	"testing"
)

func TestNewQuickFindSet(t *testing.T) {
	wantSize := 15

	set := NewQuickFindSet(wantSize)
	if set == nil {
		t.Fatalf("NewQuickFindSet(%d) = nil, want something != nil", wantSize)
	}

	if got := set.Size(); got != wantSize {
		t.Errorf("NewQuickFindSet(%d) = %d, want %d", wantSize, got, wantSize)
	}
}

func TestQuickFindSet_Union(t *testing.T) {
	tests := []struct {
		name string
		set  QuickFindSet
		a, b int
		want QuickFindSet
	}{
		{"Connect 0 and 1", QuickFindSet{[]int{0, 1, 2, 3, 4, 5, 6, 7}}, 0, 1, QuickFindSet{[]int{1, 1, 2, 3, 4, 5, 6, 7}}},
		{"Connect 1 and 0", QuickFindSet{[]int{0, 1, 2, 3, 4, 5, 6, 7}}, 1, 0, QuickFindSet{[]int{0, 0, 2, 3, 4, 5, 6, 7}}},
		{"Connect 3 and 4", QuickFindSet{[]int{0, 1, 2, 3, 4, 5, 6, 7}}, 3, 4, QuickFindSet{[]int{0, 1, 2, 4, 4, 5, 6, 7}}},
		{"Connect 0 and 7", QuickFindSet{[]int{0, 1, 2, 3, 4, 5, 6, 7}}, 0, 7, QuickFindSet{[]int{7, 1, 2, 3, 4, 5, 6, 7}}},
		{"Connect already connected", QuickFindSet{[]int{0, 1, 2, 2, 4, 5, 6, 7}}, 2, 3, QuickFindSet{[]int{0, 1, 2, 2, 4, 5, 6, 7}}},
		{"Connect two components", QuickFindSet{[]int{0, 0, 3, 0, 4, 4, 5, 3}}, 3, 2, QuickFindSet{[]int{3, 3, 3, 3, 4, 4, 5, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Clone()
			got.Union(tt.a, tt.b)
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("QuickFindSet.Union(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestQuickFindSet_IsConnected(t *testing.T) {
	tests := []struct {
		name string
		set  QuickFindSet
		a, b int
		want bool
	}{
		{"0 and 1 connected", QuickFindSet{[]int{3, 3, 0, 5, 0, 5, 5, 5}}, 0, 1, true},
		{"2 and 4 connected", QuickFindSet{[]int{3, 3, 0, 5, 0, 5, 5, 5}}, 2, 4, true},
		{"6 and 7 connected", QuickFindSet{[]int{3, 3, 0, 5, 0, 5, 5, 5}}, 6, 7, true},
		{"0 and 2 not connected", QuickFindSet{[]int{3, 3, 0, 5, 0, 5, 5, 5}}, 0, 2, false},
		{"2 and 7 not connected", QuickFindSet{[]int{3, 3, 0, 5, 0, 5, 5, 5}}, 2, 7, false},
		{"0 and 7 not connected", QuickFindSet{[]int{3, 3, 0, 5, 0, 5, 5, 5}}, 0, 7, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.IsConnected(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("QuickFindSet.IsConnected(%v, %v, %v) = %v, want %v", tt.set, tt.a, tt.b, got, tt.want)
			}
		})
	}
}
