package main

import (
	"fmt"
)

func main() {
	var array []int
	array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(array)
	slice1 := array[0:3]
	fmt.Println(slice1)
	fmt.Println("----------------------")
	printSlice(slice1)
	//扩充，其实填充的内容是底层
	slice2 := slice1[1:13]
	printSlice(slice2)
}

func printSlice(s []int) {
	fmt.Printf("content:%v length:%d capacity:%d\n", s, len(s), cap(s))
}
