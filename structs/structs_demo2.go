package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
}

func (self *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v, %T", self, self,self)
}

func main() {
	//m := Manager{User{1, "Tom"}}
	//
	//fmt.Printf("Manager: %p\n", &m)
	//fmt.Println(m.ToString())

	s1 := [...]int{1,2,3}
	s2 := []int{1,2,3}
	s3 := make([]int,10)

	typeS1 := fmt.Sprintf("%T",s1)
	typeS2 := fmt.Sprintf("%T",s2)
	typeS3 := fmt.Sprintf("%T",s3)
	fmt.Println(typeS1,typeS2,typeS1 == typeS2)
	fmt.Println(typeS1,typeS3,typeS1 == typeS3)
	fmt.Println(typeS2,typeS3,typeS2 == typeS3)
}
