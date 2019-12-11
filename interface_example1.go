package main

import "fmt"

// Ref 1 : http://golang.site/go/article/18-Go-%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4
// This example is about empty interface를
// empty interface contains all tpyes of struct의

type Shape interface { // 자바처럼 무조건 area(), perimeter를 구현해야한다.
} //  interface를 구현하는 클래스에 해당하는 것이 아니라 모든 구조체에 적용할 수 있다.

//Rect 정의
type Rect struct {
	width, height float64
}

//Circle 정의
type Circle struct {
	radius float64
}

func showArea(shapes ...Shape) { //interface를 구현한 struct의 인스턴스들을 같은 interface의 리스트(?)로 묶일 수 있다. 거기에서 구현된 함수를 실행 할 수 있다.
	for idx, s := range shapes {
		println(idx, s)
	}
}

func printIt(v interface{}) {
	fmt.Println(v) //Tom
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)

	var x interface{}
	x = 3.14
	printIt(x)
	x = "Tom"
	printIt(x)

	var a interface{} = 3

	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1 a.(type)은 go의 type casting

	println(i) // 포인터주소 출력
	println(j) // 1 출력
}
