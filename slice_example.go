package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	// s := make([]int, 5) == s := make([]int, 5, 5)
	println(len(s), cap(s)) // len 5, cap 10
	s = []int{0, 1, 2, 3, 4, 5}
	s = s[2:5]     // 2, 3, 4
	s = s[1:]      // 3, 4
	fmt.Println(s) // 3, 4 출력

	s = append(s, 2) // 0, 1, 2 // This function is similar to python list's append
	fmt.Println(s)
	// 복수 요소들 확장
	s = append(s, 3, 4, 5) // 0,1,2,3,4,5
	fmt.Println(s)
	var arr []int
	// nill works like null
	if arr == nil {
		println("Nil Slice")
	}
	println(len(arr), cap(arr)) // 모두 0

	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	// cap가 부족하면 2배를 추가로 할당한다 C++의 벡터와 유사하다.
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력
}
