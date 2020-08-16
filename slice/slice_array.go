package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	slice := array[0:3]
	fmt.Println(slice)
	fmt.Printf("%p %p", &array, &slice) //切片与数组不共用地址

}
