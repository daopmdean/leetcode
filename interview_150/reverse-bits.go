package interview_150

import "strconv"

func reverseBits(num uint32) uint32 {
	str := uint32ToString(num)
	str = reverseStr(str)
	return stringToUint32(str)
}

func uint32ToString(num uint32) string {
	str := strconv.FormatUint(uint64(num), 2)
	for len(str) < 32 {
		str = "0" + str
	}
	return str
}

func stringToUint32(str string) uint32 {
	u, _ := strconv.ParseUint(str, 2, 32)
	return uint32(u)
}
