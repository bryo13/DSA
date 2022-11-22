package bubble

import (
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("acending", func(t *testing.T) {
		array := [5]int16{4, 3, 2, 1, 0}
		got := SortArray(array)
		want := [5]int16{0, 1, 2, 3, 4}

		if got != want {
			t.Errorf("want %v but got %v ", want, got)
		}
	})
}
