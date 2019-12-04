package main

func main() {

	sum := 0 // int declaration

	for i := 1; i <= 100; i++ { // prefix operation is not suportted in Go.
		sum += 1 // This for statement looks like C for statment
	}

	println(sum)

	n := 1
	for n < 100 { // This statement looks like
		n *= 2
	}
	println(n)

	//for {
	// Infinite Loop
	//}

	names := []string{"Math", "Science", "English"}

	for idx, item := range names { // This statement looks like
		println(idx, item)
	}

	// brea, goto, labled of both are the same in oher language.
}
