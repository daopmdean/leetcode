package interview_150

func isValid(s string) bool {
	//"([)]{}"
	stack := []rune{}
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			if r == ')' && stack[len(stack)-1] != '(' {
				return false
			} else if r == '}' && stack[len(stack)-1] != '{' {
				return false
			} else if r == ']' && stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
