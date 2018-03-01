package goalgorithms

func mergeTopDown(a []int, b []int, i, size int) {
	l := i
	lsize := size/2 + size%2
	r := i + lsize
	rsize := size - lsize

	if lsize > 1 {
		mergeTopDown(a, b, l, lsize)
	}
	if rsize > 1 {
		mergeTopDown(a, b, r, rsize)
	}

	lmax := l + lsize
	rmax := r + rsize

	z := 0
	for z < size {
		if l == lmax {
			b[z] = a[r]
			r++
		} else if r == rmax {
			b[z] = a[l]
			l++
		} else if a[l] <= a[r] {
			b[z] = a[l]
			l++
		} else {
			b[z] = a[r]
			r++
		}
		z++
	}

	for z := 0; z < size; z++ {
		a[i+z] = b[z]
	}
}

// MergeSortTopDown performs in-place sort of int slice in ascending order.
func MergeSortTopDown(a []int) {
	b := make([]int, len(a))
	mergeTopDown(a, b, 0, len(a))
}

func mergeTopDown2(a []int, b []int, left, right int) {
	middle := left + ((right - left) / 2)

	if middle-left > 1 {
		mergeTopDown2(a, b, left, middle)
	}
	if right-middle > 1 {
		mergeTopDown2(a, b, middle, right)
	}

	s := right - left
	l := left
	r := middle
	z := 0
	for z < s {
		if l == middle {
			b[z] = a[r]
			r++
		} else if r == right {
			b[z] = a[l]
			l++
		} else if a[l] <= a[r] {
			b[z] = a[l]
			l++
		} else {
			b[z] = a[r]
			r++
		}
		z++
	}

	for s := 0; s < z; s++ {
		a[left+s] = b[s]
	}
}

// MergeSortTopDown2 performs in-place sort of int slice in ascending order.
func MergeSortTopDown2(a []int) {
	b := make([]int, len(a))
	mergeTopDown2(a, b, 0, len(a))
}
