package debugutil

import "fmt"

func PrintSlice(s []int) {
	fmt.Printf("content:%v length:%d capacity:%d\n", s, len(s), cap(s))
}
