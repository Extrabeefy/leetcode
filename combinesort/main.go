package main

import "log"

// ====================== Problem ================
// Example 3: Given two sorted integer arrays arr1 and arr2, return a new array that combines both of them and is also sorted.

func main() {
	arr1, arr2 := []int{1, 4, 7, 20}, []int{3, 5, 6}
	log.Print(combineArray(arr1, arr2))
}

func combineArray(arr1 []int, arr2 []int) []int {
	ans := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			ans = append(ans, arr1[i])
			i++
		} else {
			ans = append(ans, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		ans = append(ans, arr1[i])
		i++
	}

	for i < len(arr2) {
		ans = append(ans, arr2[j])
		j++
	}

	return ans
}
