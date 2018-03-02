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
	b := make([]int, len(a), len(a))
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
	b := make([]int, len(a), len(a))
	mergeTopDown2(a, b, 0, len(a))
}

func mergeTopDown3(a []int, b []int, left, right int) {
	middle := left + ((right - left) / 2)

	if middle-left > 1 {
		mergeTopDown3(a, b, left, middle)
	}
	if right-middle > 1 {
		mergeTopDown3(a, b, middle, right)
	}

	l := left
	r := middle
	for z := left; z < right; z++ {
		if l < middle && (r == right || a[l] <= a[r]) {
			b[z] = a[l]
			l++
		} else {
			b[z] = a[r]
			r++
		}
	}

	for z := left; z < right; z++ {
		a[z] = b[z]
	}
}

// MergeSortTopDown3 performs in-place sort of int slice in ascending order.
func MergeSortTopDown3(a []int) {
	b := make([]int, len(a), len(a))
	mergeTopDown3(a, b, 0, len(a))
}

// MergeSortBottomUp1 performs in-place sort of int slice in ascending order.
func MergeSortBottomUp1(a []int) {
	b := make([]int, len(a), len(a))
	s := 1
	for s < len(a) {
		for left, right := 0, s; left < len(a); left, right = left+s*2, right+s*2 {
			z := 0
			l := left
			ls := l + s
			if ls > len(a) {
				ls = len(a)
			}
			r := right
			rs := r + s
			if rs > len(a) {
				rs = len(a)
			}
			for l < ls || r < rs {
				if l < ls && (r >= rs || a[l] <= a[r]) {
					b[z] = a[l]
					l++
				} else {
					b[z] = a[r]
					r++
				}
				z++
			}
			for m := 0; m < z; m++ {
				a[left+m] = b[m]
			}
		}
		s *= 2
	}
}
