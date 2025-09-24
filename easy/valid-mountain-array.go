package easy

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	up, down := false, false
	for i, v := range arr {
		if i == 0 {
			continue
		}
		if i == 1 && v <= arr[0] {
			return false
		}
		if v == arr[i-1] {
			return false
		} else if v > arr[i-1] {
			up = true
			if down {
				return false
			}
		} else {
			down = true
		}
	}

	if up && down {
		return true
	}

	return false
}
