package main

import "fmt"

func main() {
	var i,j,k int = 5,6,7
	voidRun(i,j,k)
}

func voidRun(o ...interface{}) {
	for _,v := range o{
		fmt.Printf("%T\n",v)
	}
}
