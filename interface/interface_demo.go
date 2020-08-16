package main

import "fmt"

type creature interface {
	eat() string //这里其实就是函数的声明部分
}

type human struct {
	foodType string
}

func main() {
	var newCreature creature = human{"rice"}
	fmt.Println(newCreature.eat())
}

/*
	补充接口里的函数声明
	(h human) 表明human实现了creature
	“eat() string” 就像是充当了标识符，谁实现了 “eat() string” 就实现了creature

*/
func (h human) eat() string {
	fmt.Printf("%T eat %s\n", h, h.foodType)
	return h.foodType
}
