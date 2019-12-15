package main

// Ref 1 : http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90
import "fmt"

func main() {
	// 정수형 채널을 생성한다
	// chan은 channel을 의미한다.
	ch := make(chan int)

	go func() {
		// main의 변수에 접근할 수 있다.
		ch <- 123 //채널에 123을 보낸다
	}()
	// thread와 main의 속도에 따라서 i에 123이 전달되지 않을 수도 있다?
	var i int
	i = <-ch // 채널로부터 123을  여기서 데이터를 받을때까지 대기한다 bloclking time.sleep 등이 필요하지 않다.
	println(i)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	// 위의 Go루틴이 끝날 때까지 대기
	<-done // 수신은 받는 변수 없이도 가능하다
}
