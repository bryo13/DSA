package list

import (
	"testing"
)

func TestList(t *testing.T) {

	t.Run("singly", func(t *testing.T) {
		root := new(Node)
		root = nil
		AddFront(root, 11)
		AddFront(root, 12)
		list := Iterate(root)
		got := list[0]
		want := 11

		if got != want {
			t.Errorf("wanted %v but got %v ", want, got)
		}
	})
}
