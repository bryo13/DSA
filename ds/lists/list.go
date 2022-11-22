package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func AddFront(n *Node, k int) int {
	if root == nil {
		root = &Node{k, nil}
		return 0
	}

	if k == n.Value {
		fmt.Println("Already in list")
		return -1
	}

	if n != nil {
		n.Next = &Node{k, nil}
		return 0
	}

	return AddFront(n.Next, k)
}

func Iterate(n *Node) []int {
	var array []int
	if n == nil {
		fmt.Println("Empty")
		return array
	}

	if n.Next != nil {
		fmt.Println(n.Value)
		array = append(array, n.Value)
		n = n.Next
	}

	return array
}

func main() {
	root = nil
	AddFront(root, 12)
	AddFront(root, 14)
	fmt.Println(Iterate(root))
}
