package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

//类似java toString
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)


	fmt.Println("---------------")
	i:= 0
	j:=string(i)
	fmt.Println(j)
	fmt.Println("---------------")
	fmt.Println(time.Now())
}
