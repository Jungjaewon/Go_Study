package main

func main() {

	print("Variavles examples\n")
	// In Go, all uninitialized vars are 0, flase, ""
	var a int    // a value is 0
	_ = a        // a is unused
	var s string // s is ""
	_ = s
	var flag bool // flag is false
	_ = flag
	var f float64 = 11
	_ = f

	a = 10
	f = 12.0

	var i, j, k int
	_, _, _ = i, j, k
	var x, y, z int = 1, 2, 3
	_, _, _ = x, y, z

	// Short assignment ; only possible in func
	q := 1
	w := 1.1
	str := "Example"
	_, _, _ = q, w, str

	const const_i int = 10
	const const_str string = "const_exmaple"
	// const const_i int = 10
	// const const_str string = "const_exmaple"
	println(const_i)
	println(const_str)
}
