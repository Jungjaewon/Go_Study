package main

import "fmt"

// Ref 1 : http://golang.site/go/article/14-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Map

func main() {

	var idMap map[int]string
	_ = idMap
	// idMap = make(map[int]string) // 위와 동일

	tickets := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}

	val, exists := tickets["MSFT"] // 는 두가지로 해석이 될 수 있다.
	println(val)
	println(tickets["MSFT"])
	if !exists {
		println("No MSFT ticker")
	}

	var m map[int]string

	m = make(map[int]string)
	//추가 혹은 갱신
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"
	// 키에 대한 값 읽기
	str := m[134]
	println(str)
	noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
	println(noData)
	// 삭제
	delete(m, 777) // delete(map_var, key)

	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}

}
