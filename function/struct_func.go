package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) move1() {
	v.x = math.Pow(v.x, 2)
	v.y = math.Pow(v.y, 2)
}

func move2(v vertex) {
	v.x = math.Pow(v.x, 2)
	v.y = math.Pow(v.y, 2)
}

func (v *vertex) move3() {
	v.x = math.Pow(v.x, 2)
	v.y = math.Pow(v.y, 2)
}

func move4(v *vertex) {
	v.x = math.Pow(v.x, 2)
	v.y = math.Pow(v.y, 2)
}

func main() {
	newVertex := vertex{3, 3}
	fmt.Println("origin: ", newVertex) //origin:  {3 3}

	newVertex.move1()                     //处理副本所以不变
	fmt.Println("run move1: ", newVertex) //run move1:  {3 3} 没有被改动

	move2(newVertex)                      //处理副本，所以不变
	fmt.Println("run move2: ", newVertex) //run move1:  {3 3} 没有被改动

	newVertex.move3()                     //处理原件 所以改变
	fmt.Println("run move3: ", newVertex) //run move3:  {9 9} 被改动

	move4(&newVertex)                     //处理原件所以改变
	fmt.Println("run move4: ", newVertex) //	run move4:  {81 81} 被改动

	(&newVertex).move1()                  //尽管传来的是原件，但是处理的是副本
	fmt.Println("run move1: ", newVertex) //	run move4:  {81 81} 没有被改动

	/*
		由此可见，与传来的是否是原件并不重要，重要的是函数处理的对象是谁
	*/
}
