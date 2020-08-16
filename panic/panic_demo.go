package main

import (
	"fmt"
)

func main() {
	//test()
	//test2()
	//test4(0, 0)
	defer func(){
		if eil:= recover();eil!=nil{
			fmt.Println("it's ok")
		}
	}()
	x,y:=1,0
	m:=x/y
	fmt.Println(m)
}

func test() {

}

func test2() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func test3() {
	defer recover()              // 无效！
	defer fmt.Println(recover()) // 无效！
	defer func() {
		func() {
			println("defer inner")
			recover() // 无效！
		}()
	}()

	panic("test panic")
}

func test4(x, y float64) {
	var z float64

	defer func() {
		if recover() != nil {
			z = 0
		}
		fmt.Println("x / y =", z)
	}()

	z = x / y

}

func test5(){
	defer func(){
		if eil:= recover();eil!=nil{
			fmt.Println("it's ok")
		}
	}()
	x,y:=1,0
	m:=x/y
	fmt.Println(m)
}
