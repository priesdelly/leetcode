package __two_sum_test

import "testing"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for indexNum, num := range nums {

		// Find value in map
		// If target value - num value has exists in map
		// It's mean value is match answer
		// Just return index of target value and index of num value
		if indexMap, found := numMap[target-num]; found {
			return []int{indexMap, indexNum}
		}
		numMap[num] = indexNum
	}
	return nil
}

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)
	if res[0] != 0 || res[1] != 1 {
		t.Errorf("Expected [0, 1], got %v", res)
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	res := twoSum(nums, target)
	if res[0] != 1 || res[1] != 2 {
		t.Errorf("Expected [1, 2], got %v", res)
	}
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	res := twoSum(nums, target)
	if res[0] != 0 || res[1] != 1 {
		t.Errorf("Expected [0, 1], got %v", res)
	}
}

func TestTwoSum4(t *testing.T) {
	nums := []int{3, 2, 3}
	target := 6
	res := twoSum(nums, target)
	if res[0] != 0 || res[1] != 2 {
		t.Errorf("Expected [0, 2], got %v", res)
	}
}

func TestTwoSum5(t *testing.T) {
	nums := []int{1, 9, 8, 2, 3, 3}
	target := 6
	res := twoSum(nums, target)
	if res[0] != 4 || res[1] != 5 {
		t.Errorf("Expected [4, 5], got %v", res)
	}
}
