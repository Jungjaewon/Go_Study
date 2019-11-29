package main

// Ref 1 : http://golang.site/go/article/5-Go-%EB%8D%B0%EC%9D%B4%ED%83%80-%ED%83%80%EC%9E%85
func main() {
	var i int = 100
	var u uint = uint(i) // uint(i) is necessary.
	var f float32 = float32(i)
	println(f, u)
	// Go does not provide implicit type conversion
	str := "ABC"
	s := "A"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
	//byte := []byte(s) This state does not work
	//println(byte)
}
