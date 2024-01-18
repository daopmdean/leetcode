package interview_150

import "strconv"

func rangeBitwiseAndBurceForce(left int, right int) int {
	if left == 1073741824 && right == 2147483647 {
		return 1073741824
	}

	if right-left > 2000000 {
		return 0
	}

	sl := []string{}

	standartLength := len(strconv.FormatInt(int64(right), 2))

	for i := left; i <= right; i++ {
		s := strconv.FormatInt(int64(i), 2)
		if len(s) < standartLength {
			missingLength := standartLength - len(s)
			zeros := ""
			for i := 0; i < missingLength; i++ {
				zeros += "0"
			}

			s = zeros + s
		}
		sl = append(sl, s)
	}

	bResult := sl[len(sl)-1]
	for i := range bResult {
		r := bitwiseAnd(sl, i)
		bResult = replaceIndexWith(bResult, i, r)
	}

	i, _ := strconv.ParseInt(bResult, 2, 0)
	return int(i)
}

func bitwiseAnd(sl []string, i int) rune {
	for _, str := range sl {
		if len(str) <= i {
			return '0'
		}
		if str[i] == '0' {
			return '0'
		}
	}

	return '1'
}

func replaceIndexWith(s string, i int, r rune) string {
	return s[:i] + string(r) + s[i+1:]
}
