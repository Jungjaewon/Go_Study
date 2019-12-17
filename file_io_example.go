package main

// Ref 1 : http://golang.site/go/applications
import (
	"io"
	"os"
)

func main() {
	// 입력파일 열기
	fi, err := os.Open("C:\\temp\\1.txt") // only read mode
	if err != nil {
		panic(err)
	}
	defer fi.Close() // this statement will be excuted the end of main.

	// 출력파일 생성
	fo, err := os.Create("C:\\temp\\2.txt") // create modeS
	if err != nil {
		panic(err)
	}
	defer fo.Close() // this statement will be excuted the end of main.

	buff := make([]byte, 1024)

	// 루프
	for {
		// 읽기
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		println(cnt)
		// 끝이면 루프 종료
		if cnt == 0 {
			break
		}

		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}
