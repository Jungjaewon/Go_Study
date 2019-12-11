package main

// Ref 1 : http://golang.site/go/article/18-Go-%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4

import "math"

type Shape interface { // 자바처럼 무조건 area(), perimeter를 구현해야한다.
	area() float64
	perimeter() float64
} //  interface를 구현하는 클래스에 해당하는 것이 아니라 모든 구조체에 적용할 수 있다.

//Rect 정의
type Rect struct {
	width, height float64
}

//Circle 정의
type Circle struct {
	radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) { //interface를 구현한 struct의 인스턴스들을 같은 interface의 리스트(?)로 묶일 수 있다. 거기에서 구현된 함수를 실행 할 수 있다.
	for _, s := range shapes {
		a := s.perimeter() //인터페이스 메서드 호출
		println(a)
	}
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}
