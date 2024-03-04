package interview_150

func addBinary(a string, b string) string {
	additionResult := ""
	if len(a) < len(b) {
		a, b = reverseStr(b), reverseStr(a)
	} else {
		a, b = reverseStr(a), reverseStr(b)
	}

	hold := 0
	for i := range b {
		if a[i] == '1' && b[i] == '1' {
			if hold == 1 {
				additionResult += "1"
			} else {
				additionResult += "0"
				hold = 1
			}
		} else if a[i] == '0' && b[i] == '0' {
			if hold == 1 {
				additionResult += "1"
			} else {
				additionResult += "0"
			}
			hold = 0
		} else {
			if hold == 1 {
				additionResult += "0"
			} else {
				additionResult += "1"
			}
		}
	}

	for i := len(b); i < len(a); i++ {
		if a[i] == '0' {
			if hold == 1 {
				additionResult += "1"
				hold = 0
			} else {
				additionResult += "0"
			}
		} else {
			if hold == 1 {
				additionResult += "0"
			} else {
				additionResult += "1"
			}
		}
	}

	if hold == 1 {
		additionResult += "1"
	}

	return reverseStr(additionResult)
}

func reverseStr(str string) string {
	result := ""

	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}

	return result
}

/* brute force solution
func addBinary(a string, b string) string {
	x, y := stringBinaryToInt(a), stringBinaryToInt(b)
	return intToStringBinary(x + y)
}

func stringBinaryToInt(str string) int {
	result, start := 0, float64(0)

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '1' {
			result += int(math.Pow(2, start))
		}

		start++
	}

	return result
}

func intToStringBinary(i int) string {
	return strconv.FormatInt(int64(i), 2)
}
*/
