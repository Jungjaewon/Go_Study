// Ref : 1 http://golang.site/go/article/21-Go-%EB%A3%A8%ED%8B%B4-goroutine

package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 4개의 CPU 사용
	runtime.GOMAXPROCS(4)
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)
	//...
}
