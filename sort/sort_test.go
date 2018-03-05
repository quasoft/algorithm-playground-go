package goalgorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	implementations := []struct {
		name string
		sort func([]int)
	}{
		{"InsertionSortSwap", InsertionSortSwap},
		{"InsertionSortSwapOnce", InsertionSortSwapOnce},
		{"InsertionSortShift", InsertionSortShift},
		{"SelectionSort", SelectionSort},
		{"SelectionSortTemp", SelectionSortTemp},
		{"BubbleSort", BubbleSort},
		{"BubbleSortTwoLoops", BubbleSortTwoLoops},
		{"MergeSortTopDown", MergeSortTopDown},
		{"MergeSortTopDown2", MergeSortTopDown2},
		{"MergeSortTopDown3", MergeSortTopDown3},
		{"MergeSortBottomUp1", MergeSortBottomUp1},
		{"MergeSortBottomUp2", MergeSortBottomUp2},
		{"QuickSortHoare", QuickSortHoare},
		{"QuickSortHoareM3", QuickSortHoareM3},
		{"QuickSortLomuto", QuickSortLomuto},
	}
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"Mixed", []int{1, 3, 4, 5, 2, 9, 8, 0}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Already sorted", []int{0, 1, 2, 3, 4, 5, 8, 9}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Almost sorted", []int{0, 1, 2, 3, 4, 5, 9, 8}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Reversed", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"Large set of 50 values",
			[]int{
				703, 741, 755, 275, 283, 800, 120, 902, 744, 848, 473, 529, 277, 449, 172, 141, 773, 746, 308, 103,
				263, 787, 11, 259, 253, 211, 569, 613, 110, 990, 664, 588, 434, 600, 930, 145, 188, 293, 896, 719,
				534, 721, 23, 476, 671, 763, 254, 123, 838, 208,
			},
			[]int{
				11, 23, 103, 110, 120, 123, 141, 145, 172, 188, 208, 211, 253, 254, 259, 263, 275, 277, 283, 293,
				308, 434, 449, 473, 476, 529, 534, 569, 588, 600, 613, 664, 671, 703, 719, 721, 741, 744, 746, 755,
				763, 773, 787, 800, 838, 848, 896, 902, 930, 990,
			},
		},
	}
	for _, impl := range implementations {
		for _, tt := range tests {
			tosort := make([]int, len(tt.list))
			copy(tosort, tt.list)
			t.Run(tt.name, func(t *testing.T) {
				impl.sort(tosort)
				got := tosort
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("%s(%v) = %v, want %v", impl.name, tt.list, got, tt.want)
				}
			})
		}
	}
}

func BenchmarkSort(b *testing.B) {
	tests := [][]int{
		[]int{4, 1, 9, 6, 2, 5, 3, 7, 8, 0},
		[]int{
			703, 741, 755, 275, 283, 800, 120, 902, 744, 848, 473, 529, 277, 449, 172, 141, 773, 746, 308, 103,
			263, 787, 11, 259, 253, 211, 569, 613, 110, 990, 664, 588, 434, 600, 930, 145, 188, 293, 896, 719,
			534, 721, 23, 476, 671, 763, 254, 123, 838, 208,
		},
		[]int{
			728, 837, 700, 98, 488, 765, 186, 855, 362, 853, 450, 155, 629, 750, 610, 230, 650, 255, 95, 365, 494,
			680, 226, 731, 482, 711, 406, 994, 593, 161, 748, 557, 753, 942, 500, 709, 983, 990, 179, 805, 854, 335,
			877, 845, 273, 691, 839, 601, 940, 829, 851, 383, 309, 6, 24, 852, 846, 393, 505, 635, 112, 695, 901, 857,
			46, 174, 361, 956, 796, 104, 533, 865, 8, 982, 5, 832, 848, 328, 791, 291, 418, 94, 316, 512, 157, 975, 461,
			670, 998, 212, 830, 126, 400, 194, 340, 375, 920, 546, 214, 950, 874, 768, 436, 455, 444, 1000, 144, 762,
			817, 311, 726, 804, 932, 146, 183, 736, 699, 677, 345, 352, 518, 371, 129, 905, 576, 948, 399, 171, 251, 526,
			192, 120, 780, 844, 491, 355, 53, 198, 784, 158, 952, 759, 922, 730, 443, 156, 114, 814, 618, 256, 469, 132,
			170, 733, 963, 304, 898, 685, 280, 628, 807, 751, 457, 299, 127, 387, 138, 797, 27, 782, 58, 890, 929, 409,
			701, 785, 941, 575, 744, 266, 438, 326, 275, 22, 295, 283, 279, 468, 752, 113, 630, 594, 870, 842, 434, 201,
			758, 926, 288, 228,
		},
		[]int{
			11, 23, 103, 110, 120, 123, 141, 145, 172, 188, 208, 211, 253, 254, 259, 263, 275, 277, 283, 293,
			308, 434, 449, 473, 476, 529, 534, 569, 588, 600, 613, 664, 671, 703, 719, 721, 741, 744, 746, 755,
			763, 773, 787, 800, 838, 848, 896, 902, 930, 990,
		},
	}

	implementations := []struct {
		name string
		sort func([]int)
	}{
		{"InsertionSortSwap", InsertionSortSwap},
		{"InsertionSortSwapOnce", InsertionSortSwapOnce},
		{"InsertionSortShift", InsertionSortShift},
		{"SelectionSort", SelectionSort},
		{"SelectionSortTemp", SelectionSortTemp},
		{"BubbleSort", BubbleSort},
		{"BubbleSortTwoLoops", BubbleSortTwoLoops},
		{"MergeSortTopDown", MergeSortTopDown},
		{"MergeSortTopDown2", MergeSortTopDown2},
		{"MergeSortTopDown3", MergeSortTopDown3},
		{"MergeSortBottomUp1", MergeSortBottomUp1},
		{"MergeSortBottomUp2", MergeSortBottomUp2},
		{"QuickSortHoare", QuickSortHoare},
		{"QuickSortHoareM3", QuickSortHoareM3},
		{"QuickSortLomuto", QuickSortLomuto},
	}
	for _, tt := range tests {
		for _, impl := range implementations {
			b.Run(impl.name+fmt.Sprintf("%d", len(tt)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					tosort := make([]int, len(tt))
					copy(tosort, tt)
					impl.sort(tosort)
				}
			})
		}
	}
}
