package main

import "fmt"

type st struct {
	x int
}

func main() {
	a := st{1}
	b := st{2}
	fmt.Println(a==b,&a == &b)

	c := [...]int{1,2,3}
	d := [...]int{1,2,3}
	fmt.Println(c==d,&c==&d)
	fmt.Println("----------")
	var s [3]int
	fmt.Println(s)


	//data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	//s := data[8:] //[8 9]
	//s2 := data[:5]// [0 1 2 3 4]
	//
	//copyLen := copy(s2, s) // dst:s2, src:s 从s往s2复制，复制长度为2
	//
	//fmt.Println(copyLen)
	//fmt.Println(s)
	//fmt.Println(s2)
	//fmt.Println(data)

	//data := [...]int{1,2,3,4,5}
	//s1 := data[:]
	//data[2] = 100
	//fmt.Println(s1)

	//data := [...]int{1,2,3,4,10:1}
	//s:=data[:2:3]
	//s2 := append(s,100,200)
	//fmt.Println(&s[0],&s2[0])

	//s1 := make([]int, 0, 1)
	//s1 = append(s1,1)
	//fmt.Printf("%p\n", &s1[0])
	//
	//s2 := append(s1, 2)
	//fmt.Printf("%p\n", &s2[0])
	//
	//fmt.Println(s1, s2)

	//data := [...]int{0,1,2,3,4,5,6,7,8,9}
	//
	//s1 := data[:3]
	//s2 := append(s1,100,200)
	//s1 = append(s1,400)
	//fmt.Println(data)
	//fmt.Println(s1)
	//fmt.Println(s2)

	//s1 := []int{1,2,3}
	////s1 := make([]int,100,100)
	////s1 = append(s1,1,2,3)
	//fmt.Printf("%p\n", &s1)
	//
	//s2 := append(s1, 4)
	//fmt.Printf("%p\n", &s2)
	//
	//fmt.Println(s1[2],s2[2],&s1[2], &s2[2])

}
