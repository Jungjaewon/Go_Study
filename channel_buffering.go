package main

// Ref 1: http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90

import "fmt"

func main() {
	ch := make(chan int, 1)
	//ch := make(chan int) // unbuffered channel
	//수신자가 없더라도 보낼 수 있다.
	ch <- 101
	var i int
	i = <-ch // i <- ch
	fmt.Println(i)
}
