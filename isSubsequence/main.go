package main

import "log"

// ====================== Problem =========================
// Example 4: 392. Is Subsequence.

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

// A subsequence of a string is a sequence of characters that can be obtained by deleting some (or none) of the characters from the original string, while maintaining the relative order of the remaining characters. For example, "ace" is a subsequence of "abcde" while "aec" is not.

func main() {
	s1 := "ace"
	s2 := "abcde"
	// s2 := "gdfse"
	b1 := []byte(s1)
	b2 := []byte(s2)

	log.Print(isSubsequence(b1, b2))
}

func isSubsequence(s []byte, t []byte) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	return i == len(s)
}
