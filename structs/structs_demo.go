package main

import "fmt"

type cat struct {
	name string
	age  int
}

func main() {
	newCat := cat{age: 1,name: "huahua"}
	newCat.printInstanceCat()
	newCat.printRealCat()
	fmt.Println(newCat)
	newCat = cat{}

	//var a = []int{1, 2, 3, 4, 5}
	////a := []int{1,2,3}
	//a = append(a, 6, 7, 8, 9, 10)
	//a[8] = 6
	//fmt.Println(a)
	//fmt.Println("------------------------")
	//for k, v := range a {
	//	fmt.Println(k, v)
	//}
	//fmt.Println("------------------------")
	//
	//newCat1 := cat{name: "huahua", age: 5}
	//p := &newCat1
	//p.age = 8
	//newCat2 := cat{"miaomiao", 3}
	//var q *cat = &newCat2
	//q.name = "cute"
	//fmt.Println(newCat1, newCat2)
}

func (c cat) printInstanceCat(){
	fmt.Printf("%p\n",&c)
	c.name = "duoduo"
	fmt.Println(c)
}

func (c *cat) printRealCat(){
	fmt.Printf("%p\n",&c)
	c.name = "duoduo"
	fmt.Println(c)
}