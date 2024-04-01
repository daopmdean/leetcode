package interview_150

import "sort"

func longestCommonPrefix(strs []string) string {
	common := ""

	sort.Strings(strs)

	first, last := strs[0], strs[len(strs)-1]
	if len(first) > len(last) {
		first, last = last, first
	}

	for i := 0; i < len(first); i++ {
		if first[i] != last[i] {
			return common
		}
		common += string(first[i])
	}

	return common
}
