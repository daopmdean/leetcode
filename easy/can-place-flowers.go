package easy

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		n--
	}

	skipNext := false
	for i := 0; i < len(flowerbed)-1; i++ {
		if skipNext {
			if i == len(flowerbed)-2 && flowerbed[i+1] == 0 && n == 1 {
				return true
			}
			skipNext = false
			continue
		}

		if flowerbed[i] == 1 {
			skipNext = true
			continue
		}

		if flowerbed[i+1] == 0 {
			n--
			skipNext = true
		}

		if n == 0 {
			return true
		}
	}

	if n > 0 {
		return false
	}

	return true
}
