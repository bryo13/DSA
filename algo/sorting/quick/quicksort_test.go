package quicksort

import "testing"

func TestQuicksort(t *testing.T) {
	array := [4]int{4, 3, 2, 1}
	want := [4]int{1, 2, 3, 4}

	got := QuickSort(array, 0, len(array)-1)
	if want != got {
		t.Errorf("wanted %v but got %v ", want, got)
	}
}
