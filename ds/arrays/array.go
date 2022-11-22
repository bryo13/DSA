package main

import "fmt"

type Arr []int

func main() {
	a := new(Arr)

	*a = append(*a, 23)
	*a = append(*a, 23)
	*a = append(*a, 23)
	fmt.Println(len(*a))
}
