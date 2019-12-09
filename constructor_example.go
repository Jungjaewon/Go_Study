package main

import "fmt"

// Ref 1: http://golang.site/go/article/16-Go-%EA%B5%AC%EC%A1%B0%EC%B2%B4-struct
// Go's constructor is not a constructor in other language like python, C++ and so on.
// Go's constructor looks like a help function
type dict struct {
	data map[int]string
}

type person struct {
	name string
	age  int
}

//생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}

func newPerson() *person {
	p := person{}
	p.name = "Default"
	p.age = -1
	return &p
}

func main() {
	dic := newDict() // 생성자 호출
	dic.data[1] = "A"

	person_exaple := newPerson()
	fmt.Println(person_exaple)
}
