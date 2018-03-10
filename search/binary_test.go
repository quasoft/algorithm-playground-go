package goalgorithms

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	implementations := []struct {
		name   string
		search func(needle int, haystack []int) (int, bool)
	}{
		{"BinarySearchRecursive", BinarySearchRecursive},
	}
	tests := []struct {
		name     string
		haystack []int
		needle   int
		wantIdx  int
		wantOK   bool
	}{
		{"Single existing value", []int{2}, 2, 0, true},
		{"Single non-existing value", []int{2}, 4, 0, false},
		{"First of even number of values", []int{0, 16, 25, 32, 46}, 0, 0, true},
		{"Middle of even number of values", []int{0, 16, 25, 32, 46}, 32, 3, true},
		{"Last of even number of values", []int{0, 16, 25, 32, 46}, 46, 4, true},
		{"First of odd number of values", []int{40, 41, 42, 64, 65, 75}, 40, 0, true},
		{"Middle of odd number of values", []int{40, 41, 42, 64, 65, 75}, 42, 2, true},
		{"Last of odd number of values", []int{40, 41, 42, 64, 65, 75}, 75, 5, true},
		{"Repeating values", []int{24, 24, 24, 33, 42, 42, 42}, 33, 3, true},
		{"Close non existing value", []int{24, 24, 24, 33, 42, 42, 42}, 32, 0, false},
	}
	for _, impl := range implementations {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotIdx, gotOK := impl.search(tt.needle, tt.haystack)
				if gotOK != tt.wantOK {
					t.Errorf("%s(%v, %v) = %v, want %v", impl.name, tt.needle, tt.haystack, gotOK, tt.wantOK)
				} else if gotOK {
					if gotIdx != tt.wantIdx {
						t.Errorf("%s(%v, %v) = %v, want %v", impl.name, tt.needle, tt.haystack, gotIdx, tt.wantIdx)
					}
				}
			})
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	tests := []struct {
		name     string
		haystack []int
	}{
		{"random", []int{
			11, 23, 103, 110, 120, 123, 141, 145, 172, 188, 208, 211, 253, 254, 259, 263, 275, 277, 283, 293,
			308, 434, 449, 473, 476, 529, 534, 569, 588, 600, 613, 664, 671, 703, 719, 721, 741, 744, 746, 755,
			763, 773, 787, 800, 838, 848, 896, 902, 930, 990,
		}},
	}

	implementations := []struct {
		name   string
		search func(needle int, haystack []int) (int, bool)
	}{
		{"BinarySearchRecursive", BinarySearchRecursive},
	}
	for _, tt := range tests {
		for _, impl := range implementations {
			b.Run(impl.name+fmt.Sprintf("_%s_%d", tt.name, len(tt.haystack)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					rndIdx := rand.Intn(len(tt.haystack))
					needle := tt.haystack[rndIdx]
					impl.search(needle, tt.haystack)
				}
			})
		}
	}
}
