package goalgorithms

// SelectionSort sorts an int slice in ascending order.
// Worst-case time compexity: O(n^2)
// Worst-case space compexity: O(n)
func SelectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for k := i + 1; k < len(a); k++ {
			if a[k] < a[min] {
				min = k
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

// SelectionSortTemp is variant of selection sort with temp varaible for min value.
func SelectionSortTemp(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := a[i]
		m := i
		for k := i + 1; k < len(a); k++ {
			if a[k] < min {
				min = a[k]
				m = k
			}
		}
		a[i], a[m] = min, a[i]
	}
}
