package main

// Ref 1 : http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90
import "fmt"

func sendChan(ch chan<- string) {
	ch <- "Data"
	ch <- "Test"
	// x := <-ch // 에러발생
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
	data = <-ch
	fmt.Println(data)
}

func main() {
	ch := make(chan string, 2)
	sendChan(ch)
	receiveChan(ch)
}
