package leetcode

/*
Given a string s, return the longest palindromic substring in s.
*/

func longestPalindrome(s string) string {
	stringLength := len(s)

	// If the string is empty there is no palindromic substring
	if s == "" || stringLength == 0 {
		return ""
	}

	// maxLength tracks the longest palindromic substring found
	maxLength := 1
	// low is the left edge of a substring window
	// high is the right edge of a substring window
	// start is the starting index of the longest found palindromic substring
	low, high, start := 0, 0, 0

	// loop through every character of the string
	for i := 1; i < stringLength; i++ {
		// check for an even length palindrome around i
		// an even length palindrome would have the form cbaabc

		// set the upper and lower bounds of the window
		low = i - 1
		high = i

		// check that the window is within the length of the string
		// check if the values at the low and high bounds are equal
		// while they are equal continue checking the length of this palindromic substring and expanding
		for ; low >= 0 && high < stringLength && s[low] == s[high]; {

			// equal values means the palindromic substring continues
			// check if the current substring is longer than the max found
			if high - low + 1 > maxLength {
				start = low
				maxLength = high - low + 1
			}
			// expand the window outward from i
			low--
			high++
		}

		// check for an odd length palindrome around i
		// and odd length palindrome would have the form cbabc

		// set the upper and lower bounds of the window
		low = i - 1 
		high = i + 1

		// check that the window is within the length of the string
		// check if the values at the low and high bounds are equal
		// while they are equal continue checking the length of this palindromic substring and expanding
		for ; low >= 0 && high < stringLength && s[low] == s[high]; {

			// equal values means the palindromic substring continues
			// check if the current substring is longer than the max found
			if high - low + 1 > maxLength {
				start = low
				maxLength = high - low + 1
			}

			// expand the window outward from i
			low--
			high++
		}
	}

	// return the substring of s using the start index and length of the substring
	return s[start:start+maxLength]
}