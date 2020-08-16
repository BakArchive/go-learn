package main

import "fmt"

func main() {
	var slice = make([]int, 5, 10)
	fmt.Println(slice)
	change(slice)
	fmt.Println(slice)
	slice = slice[:cap(slice)]
	fmt.Println(slice)
}

func change(slice []int) {
	if slice != nil && len(slice) > 0 {
		slice[0] = 1
	}
}
