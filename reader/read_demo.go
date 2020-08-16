package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 1)
	for {
		n, err := r.Read(b)//n 每次读取的字节数 err 是否读完
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
