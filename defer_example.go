package main

// Ref 1 : http://golang.site/go/article/20-Go-defer%EC%99%80-panic
// def가 open 뒤에 위치한다.
import "os"

func main() {
	f, err := os.Open("./src/1.txt")
	if err != nil {
		panic(err)
	}

	// main 마지막에 파일 close 실행
	defer f.Close() // defer works like finally in other lanauges.
	// defered function is called when before the function which call the defered function is ended.

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}
