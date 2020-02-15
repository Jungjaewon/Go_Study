package main

// Complete the isBalanced function below.
func isBalanced(s string) string {

	if len(s) == 1 {
		return "NO"
	}
	var stack []string
	var strLen int = len(s)

	for i := 0; i < strLen; i++ {

		c := string(s[i])
		if c == "{" || c == "(" || c == "[" {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return "NO"
		} else if c == "}" && stack[len(stack)-1] == "{" {
			stack = stack[:len(stack)-1] // stack.pop();
		} else if c == ")" && stack[len(stack)-1] == "(" {
			stack = stack[:len(stack)-1]
		} else if c == "]" && stack[len(stack)-1] == "[" {
			stack = stack[:len(stack)-1]
		} else {
			return "NO"
		}

	}

	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func mian() {

}
