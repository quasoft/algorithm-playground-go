package goalgorithms

// BubbleSort is an implementation of bubble sort with one loop.
// Sorts the int slice in-place in ascending order.
// Worst-case time compexity: O(n^2).
// Don't ever use in production. For small sets use insertion or selection sort instead.
func BubbleSort(a []int) {
	n := len(a)
	i := 1
	for i < n {
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
		}
		i++
		if i == n {
			i = 1
			n--
		}
	}
}

// BubbleSortTwoLoops is an implementation of bubble sort with two loops.
// Sorts the int slice in-place in ascending order.
func BubbleSortTwoLoops(a []int) {
	for n := len(a); n > 0; n-- {
		for i := 1; i < n; i++ {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
			}
		}
	}
}
