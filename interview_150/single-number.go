package interview_150

func singleNumber(nums []int) int {
	s := 0
	//m := map[int]bool{}

	for _, num := range nums {
		// if m[num] {
		//     s -= num
		// } else {
		//     m[num] = true
		//     s += num
		// }
		s = s ^ num
	}

	return s
}
