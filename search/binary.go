package goalgorithms

func searchRecursive(needle int, haystack []int, left, right int) (int, bool) {
	if right-left < 2 {
		return left, haystack[left] == needle
	}

	middle := left + (right-left)/2
	if needle < haystack[middle] {
		return searchRecursive(needle, haystack, left, middle)
	} else if needle > haystack[middle] {
		return searchRecursive(needle, haystack, middle, right)
	} else {
		return middle, true
	}
}

// BinarySearchRecursive performs a binary search for needle in haystack
// using recursion. Takes O(log(n)) time.
func BinarySearchRecursive(needle int, haystack []int) (int, bool) {
	return searchRecursive(needle, haystack, 0, len(haystack))
}

// BinarySearchLinear performs a binary search for needle in haystack
// using iteration. Takes O(log(n)) time.
// Slightly faster than recursive variant and uses less memory.
func BinarySearchLinear(needle int, haystack []int) (int, bool) {
	left := 0
	right := len(haystack)
	for right > left {
		middle := left + (right-left)/2

		if needle < haystack[middle] {
			right = middle
		} else if needle > haystack[middle] {
			left = middle + 1
		} else {
			return middle, true
		}
	}
	return 0, false
}
