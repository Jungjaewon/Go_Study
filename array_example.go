package main

// Ref1 :

func main() {
	var arr [3]int  //정수형 3개 요소를 갖는 배열 a 선언
	println(arr[0]) // 0
	println(arr[1]) // 0
	println(arr[2]) // 0
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	println(arr[1]) // 2 출력

	// initialization
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3} //배열크기 자동으로
	_ = arr1
	_ = arr2
	var multiArray [3][4][5]int // 정의
	multiArray[0][1][2] = 10    // 사용

	var arr3 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}
	println(arr3[1][2])
}
