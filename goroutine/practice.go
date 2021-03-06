package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	SendValue(t, ch)
	close(ch)
}

func SendValue(t *tree.Tree, ch chan int) {
	if t != nil {
		SendValue(t.Left, ch)
		ch <- t.Value
		SendValue(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v := range ch1 {
		if v != <- ch2 {
			return false
		}
	}

	return true
}

func main() {
	flag := Same(tree.New(1),tree.New(1))
	fmt.Println(flag)

	flag = Same(tree.New(1),tree.New(2))
	fmt.Println(flag)

}
