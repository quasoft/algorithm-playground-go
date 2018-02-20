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
