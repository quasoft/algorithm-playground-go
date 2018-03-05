package goalgorithms

func hoarePartition(a []int, left, right int) int {
	p := a[left+(right-left)/2]
	i := left
	j := right - 1
	for {
		for a[i] < p {
			i++
		}

		for a[j] > p {
			j--
		}

		if i >= j {
			return j
		}

		a[i], a[j] = a[j], a[i]
	}
}

func quickSortHoare(a []int, left, right int) {
	if right-left < 2 {
		return
	}
	p := hoarePartition(a, left, right)
	quickSortHoare(a, left, p)
	quickSortHoare(a, p+1, right)
}

// QuickSortHoare performs in-place sort of int slice in ascending order using Hoare partitioning.
// Worst case time compexity: O(n^2)
// Average time compexity: O(n log(n))
// Worst case space compexity: O(n)
func QuickSortHoare(a []int) {
	quickSortHoare(a, 0, len(a))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func medianOfThree(a []int, v1, v2, v3 int) int {
	if v1 > v2 && v1 < v3 || v1 > v3 && v1 < v2 {
		return v1
	} else if v2 > v1 && v2 < v3 || v2 > v3 && v2 < v1 {
		return v2
	} else {
		return v3
	}
}

func hoarePartitionM3(a []int, left, right int) int {
	p := medianOfThree(a, a[left], a[left+(right-left)/2], a[right-1])
	i := left
	j := right - 1
	for {
		for a[i] < p {
			i++
		}

		for a[j] > p {
			j--
		}

		if i >= j {
			return j
		}

		a[i], a[j] = a[j], a[i]
	}
}

func quickSortHoareM3(a []int, left, right int) {
	if right-left < 2 {
		return
	}
	p := hoarePartitionM3(a, left, right)
	quickSortHoareM3(a, left, p)
	quickSortHoareM3(a, p+1, right)
}

// QuickSortHoareM3 performs in-place sort of int slice in ascending order using Hoare
// partitioning and median of three for pivot selection.
// Worst case time compexity is still O(n^2), but not so for already sorted arrays.
// Average time compexity: O(n log(n)).
// Worst case space compexity: O(n).
func QuickSortHoareM3(a []int) {
	quickSortHoareM3(a, 0, len(a))
}

func lomutoPartition(a []int, left, right int) int {
	p := a[right-1]
	i := left
	for j := left; j < right-1; j++ {
		if a[j] < p {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[right-1] = a[right-1], a[i]
	return i
}

func quickSortLomuto(a []int, left, right int) {
	if right-left < 2 {
		return
	}
	p := lomutoPartition(a, left, right)
	quickSortLomuto(a, left, p)
	quickSortLomuto(a, p+1, right)
}

// QuickSortLomuto performs in-place sort of int slice in ascending order using Lomuto partitioning.
// Worst case time compexity: O(n^2)
// Average time compexity: O(n log(n))
// Worst case space compexity: O(n)
func QuickSortLomuto(a []int) {
	quickSortLomuto(a, 0, len(a))
}
