package main

// Ref 1 : http://golang.site/go/article/104-JSON-%EC%82%AC%EC%9A%A9
import (
	"encoding/json"
	"fmt"
)

func main() {
	// 테스트용 JSON 데이타
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	// JSON 디코딩
	var mem Member
	err := json.Unmarshal(jsonBytes, &mem) // decoding
	if err != nil {
		panic(err)
	}

	// mem 구조체 필드 엑세스
	fmt.Println(mem.Name, mem.Age, mem.Active)
}
