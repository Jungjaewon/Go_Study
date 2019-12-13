package main

// Ref 1 : http://golang.site/go/article/20-Go-defer%EC%99%80-panic
// 용어는 defer이고 핸들러를 등록하는 것과 유사하다.
// 조금 더 예제가 필요할 것 같다.
import (
	"fmt"
	"os"
)

func openFile(fn string) {
	// defere 함수. panic 호출시 실행됨
	defer func() { // What is "r"? // no arguments for r, () for arguments // similar to try-catch
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	// 파일 close 실행됨
	defer f.Close()

}

func main() {
	openFile("./src/1.txt")
	println("Done") // 이 문장 실행됨
}
