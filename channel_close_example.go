package main

// Ref 1 : http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90
func main() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch) // 채널로 송신은 안되지만 수신은 가능하다.

	// 채널 수신
	println(<-ch)
	println(<-ch)

	// := <-ch는 데이터만 반환하거나 성공여부를 반환하는 하나 혹은 두개의 반환값을 가진다.
	if _, success := <-ch; !success {
		println("더이상 데이타 없음.")
	}
}
