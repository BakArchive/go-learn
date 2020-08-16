package main

import (
	"fmt"
	"time"
)

func say(s string) {
	i := 0
	for ; ; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		if s=="world"&&i==5{
			break
		}
	}
	fmt.Println(s,":",i)
}

func main() {
	go say("world")
	go say("hello")
	say("Hello World")
}
