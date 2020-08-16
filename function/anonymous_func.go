package main

import (
	"fmt"
)


func main() {
	test()
	//f:=test2()
	//f()
}

func test(){
	i, n := 1 ,2
	defer func(a int){
		fmt.Println("defer:", a , n) //n被闭包引用
	}(i) //复制i的值
	i , n = i+1,n+2
	fmt.Println(i , n)
}

