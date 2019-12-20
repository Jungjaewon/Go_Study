package main

import (
	"encoding/json"
	"fmt"
)

// Ref 1 : http://golang.site/go/article/104-JSON-%EC%82%AC%EC%9A%A9
//Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	// Go 데이타
	mem := Member{"Alex", 10, true}
	// map or structure can be encodded to json
	// JSON 인코딩
	// json key must be string
	jsonBytes, err := json.Marshal(mem) // Marshal is encoding function
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonBytes)
	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)
}
