package main

// Ref : 1 http://golang.site/go/article/21-Go-%EB%A3%A8%ED%8B%B4-goroutine
import (
	"fmt"
	"sync"
)

func main() {
	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup // We need more example of this statement.
	wait.Add(2)             //  개수가 안맞으면 wait.Wait()에서 영원히 기다리나?

	// 익명함수를 사용한 goroutine
	// 모든함수에서 defer와 실행 순서를 바꿔도 문제가 없다.
	// 그리고 main에 선언해둔 wait에 func 함수에서 접근할 수 가 있다.
	go func() {
		fmt.Println("Hello")
		wait.Done() //끝나면 .Done() 호출
		// defer wait.Done()
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		fmt.Println(msg)
		wait.Done() //끝나면 .Done() 호출
		// defer wait.Done()
	}("Hi")

	wait.Wait() //Go루틴 모두 끝날 때까지 대기 // This statememnt looks like Thread.join()
}
