package main

// Ref 1 : http://golang.site/go/article/17-Go-%EB%A9%94%EC%84%9C%EB%93%9C
//Rect - struct 정의
type Rect struct {
	width, height int
}

//Rect의 area() 메소드
func (r Rect) area() int { // (r Rect) called receiver 이 부분이 차이점
	return r.width * r.height
}

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++ //  instance's valuable is changed.
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() //메서드 호출 // 클래스의 인스턴스에서 . 으로 메서드를 호출한다.
	println(rect.width, area)
	area = rect.area2() //메서드 호출 // 클래스의 인스턴스에서 . 으로 메서드를 호출한다.
	println(rect.width, area)
}
