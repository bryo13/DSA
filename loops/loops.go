package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("outer loop")
		for j := 0; j == i; j++ {
			fmt.Println("inner")
		}
	}
}
