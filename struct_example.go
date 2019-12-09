// Ref 1 : http://golang.site/go/article/16-Go-%EA%B5%AC%EC%A1%B0%EC%B2%B4-struct
package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}

func main() {
	// person 객체 생성
	// var p1 person
	p := person{}
	p1 := person{"Jung", 29}
	// p1 := person{name: "Jung", age: 29} // we can use field name
	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	p2 := new(person)
	p2.name = "Yi" // p가 포인터라도 . 을 사용한다
	p2.age = 30

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)
}
