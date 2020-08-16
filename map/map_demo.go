package main

type user struct {
	name string
}

func main() {
	m := map[int]user{
		1: {"user1"},
	}

	//完整替换
	m[1] = user{"user2"}


	//使用指针
	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "user2"

	p:= &user{"name"}
	(*p).name = "dd"
}
