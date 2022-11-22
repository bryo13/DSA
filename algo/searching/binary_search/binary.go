package binary_search

import (
	"sort"
)

// IterativeBinarySearch return the index of the found item or 0
func IterativeBinarySearch(a []int, n int, key int) int {
	low := 1
	high := n

	// for binary search to work, the array has to be sorted
	if !sort.IntsAreSorted(a) {
		sort.Ints(a)
	}
	for low <= high {

		mid := (low + high) / 2

		if key == a[mid] {
			return mid
		}

		if key < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	// not found
	return 0
}

func RecussiveBinarySearch(a []int, high, low, key int) int {

	// for binary search to work, the array has to be sorted
	if !sort.IntsAreSorted(a) {
		sort.Ints(a)
	}

	if low == high {
		if a[low] == key {
			return low
		} else {
			return 0
		}
	} else {
		mid := (low + high) / 2
		if key == a[mid] {
			return mid
		}
		if key < a[mid] {
			return RecussiveBinarySearch(a, mid-1, low, key)
		}

		if key > a[mid] {
			return RecussiveBinarySearch(a, high, mid-1, key)
		}
	}
	return 0
}

// O(log n)
