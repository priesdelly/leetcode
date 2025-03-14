func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	numStr := strconv.Itoa(x)

	i := len(numStr) - 1
	y := 0

	for i > 0 {
		if numStr[i] != numStr[y] {
			return false
		}
		if i == y {
			return true
		}
		i--
		y++
	}

	return true
}
