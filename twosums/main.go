package main

import "log"

// ================== PROBLEM ==============
// Example 2: Given a sorted array of unique integers and a target integer, return true if there exists a pair of numbers that sum to target, false otherwise. This problem is similar to Two Sum. (In Two Sum, the input is not sorted).

// For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13, return true because 4 + 9 = 13.

func main() {
	var nums = []int{1, 2, 4, 6, 8, 9, 14, 15}
	target := 13
	log.Print(checkForTarget(nums, target))
}

func checkForTarget(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	total := 0

	for total != target && left < right {
		if nums[left]+nums[right] == target {
			return true
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}

	return false
}
