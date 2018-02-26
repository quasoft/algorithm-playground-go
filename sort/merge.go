package goalgorithms

func merge(a []int, b []int, i, size int) {
	l := i
	lsize := size/2 + size%2
	r := i + lsize
	rsize := size - lsize

	if lsize > 1 {
		merge(a, b, l, lsize)
	}
	if rsize > 1 {
		merge(a, b, r, rsize)
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

// MergeSort performs in-place sort of int slice in ascending order.
func MergeSort(a []int) {
	b := make([]int, len(a))
	merge(a, b, 0, len(a))
}
