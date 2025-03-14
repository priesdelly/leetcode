package __palindrome_number_test

import (
	"strconv"
	"testing"
)

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

func TestPalindrome1(t *testing.T) {
	num := 121
	target := true

	res := isPalindrome(num)
	if res != target {
		t.Errorf("Expected %s, got %v", boolToString(target), boolToString(res))
	}
}

func TestPalindrome2(t *testing.T) {
	num := -121
	target := false

	res := isPalindrome(num)
	if res != target {
		t.Errorf("Expected %s, got %v", boolToString(target), boolToString(res))
	}
}

func TestPalindrome3(t *testing.T) {
	num := 12121
	target := true

	res := isPalindrome(num)
	if res != target {
		t.Errorf("Expected %s, got %v", boolToString(target), boolToString(res))
	}
}

func TestPalindrome4(t *testing.T) {
	num := 10
	target := false

	res := isPalindrome(num)
	if res != target {
		t.Errorf("Expected %s, got %v", boolToString(target), boolToString(res))
	}
}

func TestPalindrome5(t *testing.T) {
	num := 2
	target := true

	res := isPalindrome(num)
	if res != target {
		t.Errorf("Expected %s, got %v", boolToString(target), boolToString(res))
	}
}

func TestPalindrome6(t *testing.T) {
	num := 182
	target := false

	res := isPalindrome(num)
	if res != target {
		t.Errorf("Expected %s, got %v", boolToString(target), boolToString(res))
	}
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
