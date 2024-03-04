package interview_150

import (
	"math"
	"strconv"
)

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
