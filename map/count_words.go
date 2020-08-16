package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "I am learning golang golang"
	fmt.Println(countWords(testString))
}

func countWords(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int, len(words))
	for _, v := range words {
		result[v] += 1
	}
	return result
}
