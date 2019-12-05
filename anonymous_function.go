// http://golang.site/go/article/10-Go-%EC%9D%B5%EB%AA%85%ED%95%A8%EC%88%98
package main

// 일급함수 : 함수를 객체처럼 다룬다. 함수를 값으로 다룬다.
// 함수를 다른 함수의 파라미터로 넘길수 있다. 파이썬과 같음
// 함수를 리턴 할 수 있다.

type calculator func(int, int) int // C언어의 typedef와 유사함 허나 순서가 반대임 앞에가 정의하는 type
// Like C, type is used for Custom Type, inferface, struct
func sum_func(n ...int) int { // function definition
	s := 0
	for _, i := range n {
		s += i
	}
	return s
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func calc_a(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	sum := func(n ...int) int { //익명함수 정의 This way is almost same without function's name
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	//변수 add 에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

	r3 := calc_a(add, 10, 20)
	println(r3)

}
