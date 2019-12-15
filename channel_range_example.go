package main

// Ref 1 : http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90
func main() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다 // channel을 닫지 않으면 range를 사용할 수 없다?
	close(ch)

	// 방법1
	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	/*
		for {
			if i, success := <-ch; success {
				println(i)
			} else {
				break
			}
		}
	*/

	// 방법2
	// 위 표현과 동일한 채널 range 문
	for i := range ch {
		println(i)
	}
}
