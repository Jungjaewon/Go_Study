package main

// Ref 1 : http://golang.site/go/article/20-Go-defer%EC%99%80-panic
// Panic works like raise exception.
// We need more examples of Panic
// panic function requires at least one argument.
import "os"

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	// 파일 close 실행됨
	defer f.Close()
}

func A() {
	panic(nil)
	B()
	println("A function Done")
}

func B() {
	C()
	println("B function Done")
}

func C() {
	println("C function Done")
}

func main() {
	//openFile("Invalid.txt")
	println("Main Start")
	A()
	println("Main Done") //이 문장은 실행 안됨
}
