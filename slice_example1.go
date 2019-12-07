package main

// Ref 1 : http://golang.site/go/article/13-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Slice
import "fmt"

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...) //sliceA = append(sliceA, 4, 5, 6)
	// ... is necesarry to where after sliceB
	//sliceC := append(sliceA, sliceB) // This line causes an error
	fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
	//fmt.Println(sliceC)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력
}
