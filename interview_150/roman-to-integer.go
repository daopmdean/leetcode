package interview_150

func romanToInt(s string) int {
	mapV := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	v := 0
	tempV := 0
	preV := 0
	curV := 0
	for _, r := range s {
		curV = mapV[r]
		if preV == 0 {
			tempV = curV
		} else if curV < preV {
			v += tempV
			tempV = curV
		} else if curV == preV {
			tempV += curV
		} else {
			tempV = curV - tempV
			v += tempV
			tempV = 0
		}
		preV = mapV[r]
		if tempV == 0 {
			preV = 0
		}
	}

	return v + tempV
}
