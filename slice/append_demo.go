package main

import "fmt"

func main() {
	slice := make([]int, 5)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	b := []int{1, 2}
	fmt.Println(b)
}
