package main

// Ref : http://golang.site/go/article/9-Go-%ED%95%A8%EC%88%98
func message(input string) { // Call by value
	println(input)
}

func message_pointer(str_point *string) { // Call by value
	*str_point = "Changed"
	println(*str_point)
}

func string_arr(str_arr ...string) {
	for _, item := range str_arr { // underbar'_' is similar to Python.
		println(item)
	}
}

func sum(nums ...int) int { // return type is here
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sum(nums ...int) (count int, total int) { // named return paramter
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return // retrun is necessary and must be empty
}

// Go does not support function overrodding?
//  message(input string) and message(str_point *string) does not work

func main() {
	str := "Message"
	message("Hello")
	message_pointer(&str)
	string_arr("Hello", "Jaewon", "Where", "Did", "You", "Want")
}
