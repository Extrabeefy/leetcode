package main

import "log"

// =========== Problem Statement =================
// Example 1: Given a string s, return true if it is a palindrome, false otherwise.

// A string is a palindrome if it reads the same forward as backward. That means, after reversing it, it is still the same string. For example: "abcdcba", or "racecar".

func main() {
	s := "dummy"
	b := []byte(s)

	log.Print(checkIfPlindrome(b))
}

func checkIfPlindrome(s []byte) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
