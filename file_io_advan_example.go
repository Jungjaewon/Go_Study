package main

// Ref 1: https://mingrammer.com/gobyexample/writing-files/
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644) // file, data, filemode
	check(err)

	f, err := os.Create("/tmp/dat2") // f is a file pointer
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2) // n2 is the number of written data.
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3) // n3 is the number of written data.

	f.Sync() // Sync를 실행하여 안정적인 스토리지에 쓰기를 플러시합니다.

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4) // n4 is the number of written data.

	w.Flush() // Flush를 사용하여 모든 버퍼링된 작업이 라이터에 적용되었는지 확인합니다.
}
