package others

import "fmt"

const size = 40

func main() {
	for i := 1; i <= size; i++ {
		fmt.Println(fib(i))
	}
	fib2()
}

func fib(num int) int {
	if num == 1 || num == 2 {
		return num - 1
	}
	return fib(num-1) + fib(num-2)
}

func fib2() {
	fibArr := [size]int{0, 1}
	for i := 2; i < size; i++ {
		fibArr[i] = fibArr[i-1] + fibArr[i-2]
	}
	fmt.Println(fibArr)
}
