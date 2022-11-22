// STEPS
// Pick a pivot - the first element in an array
// Pick the first and highest  i.e 0 and len(array) - 1
// if the first element is less than the pivot increase the index i.e first++
// if the high is greater than the pivot, decrease the high index i.e high--
// repeat the increament or decreament as long as high is larger than the first
// recussively call Quicksort on the low and high - i.e (first and result of partition) - (partion + 1 and high)
// NB
// in Go it is better to compare arrays not slices
// best case O(nlogn) worst case O(n*n)
// to improve the performance pick a pivot that isnt the first, i.e pick a random number
// or pick the middle number
package quicksort

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] < pivot {
			i++
		}

		for arr[j] >= pivot {
			j--
		}

		if i < j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	t := arr[low]
	arr[low] = arr[j]
	arr[j] = t
	return j
}

func QuickSort(a []int, low, high int) {
	if low < high {
		j := partition(a, low, high)
		QuickSort(a, low, j)
		QuickSort(a, j+1, high)
	}
	fmt.Println(a)
}
