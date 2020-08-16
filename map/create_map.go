package main

import (
	"fmt"
)

const (
	_ = iota
	SPRING
	SUMMER
	AUTUMN
	WINTER
)

func main() {
	newMap := make(map[int]string, 10)
	newMap[SPRING] = "holy shit1"
	newMap[SUMMER] = "holy shit2"
	newMap[AUTUMN] = "holy shit3"
	newMap[WINTER] = "holy shit4"
	fmt.Println(len(newMap))
	for k, v := range newMap {
		fmt.Println(k, v)
	}
}
