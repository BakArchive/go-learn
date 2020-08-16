package main

func main() {
	//ch := make(chan int) //没有缓冲区会报错
	ch := make(chan int,1)
	ch <- 4
	<- ch
}
