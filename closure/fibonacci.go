package main

import (
	"fmt"
	"math"
)

// 返回一个“返回int的函数”
func fibonacci() func() int {
	num1, num2, count := 0, 1, 0
	return func() int {
		count++
		if count == 1 {
			return num1
		} else if count == 2 {
			return num2
		} else {
			temp := num1 + num2
			num1 = num2
			num2 = temp
			return temp
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println(math.Sqrt2)
}
