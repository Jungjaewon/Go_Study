// Ref 1 : http://golang.site/go/article/4-Go-%EB%B3%80%EC%88%98%EC%99%80-%EC%83%81%EC%88%98
package main

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func nextValue2() func() (int, int) {
	i := 0
	j := 1
	return func() (int, int) {
		i++
		j++
		return i, j
	}
}

// This function works like static value in class
//

func main() {
	next := nextValue()

	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2
	println(next())        // 4

	two_next := nextValue2()

	a, b := two_next() // := means creating new vars
	println(a, b)
	a, b = two_next() // = means assignment value to variable.
	println(a, b)
	a, b = two_next()
	println(a, b)

}
