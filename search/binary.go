package goalgorithms

func searchRecursive(needle int, haystack []int, left, right int) (int, bool) {
	if right-left < 2 {
		return left, haystack[left] == needle
	}

	middle := left + (right-left)/2

	idx, ok := searchRecursive(needle, haystack, left, middle)
	if ok {
		return idx, ok
	}
	return searchRecursive(needle, haystack, middle, right)
}

// BinarySearchRecursive performs a binary search for needle in haystack
// using recursion. Takes O(log(n)) time.
func BinarySearchRecursive(needle int, haystack []int) (int, bool) {
	return searchRecursive(needle, haystack, 0, len(haystack))
}
