package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("Iterative", func(t *testing.T) {
		array := []int{5, 4, 3, 2, 1}
		key := 2
		want := 1
		got := IterativeBinarySearch(array, len(array), key)

		if want != got {
			t.Errorf("Want %v but got %v ", want, got)
		}
	})

	t.Run("Recussive", func(t *testing.T) {
		array := []int{5, 4, 3, 2, 1}
		key := 2
		want := 1
		got := RecussiveBinarySearch(array, 5, 1, key)

		if want != got {
			t.Errorf("Want %v but got %v ", want, got)
		}
	})

}
