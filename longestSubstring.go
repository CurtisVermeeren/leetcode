package leetcode

/*
Given a string s, find the length of the longest substring without repeating characters.
*/

func lengthOfLongestSubstring(s string) int {
	// position keeps track of the last index of a rune in s
	position := make(map[rune]int)

	// holds the total length of the longest substring found
	longestSubstringLength := 0
	// holds the current length of the substring being checked
	currentSubstringLength := 0
	// start indicates the starting index of the current substring
	start := 0

	// iterate the string 
	for index, char := range s {
		// check if the character was found before
		lastIndex, ok := position[char]
		
		// If the character is not in the map 
		// OR the index found is outside the current substring
		if !ok || lastIndex < index - currentSubstringLength {
			// continue the current substring and increase its length as the character is not repeated this substring
			currentSubstringLength++
		} else {
			// If the current substring is the longest then set it as longestSubstringLength
			if currentSubstringLength > longestSubstringLength {
				longestSubstringLength = currentSubstringLength
			}
			
			// update the starting position to the index of the the last character that was repeated + 1
			start = lastIndex + 1
			// adjust currentSubstringLength to accound for changing window
			currentSubstringLength = index - start + 1
		}
		// add the current character to the position tracking with its current index
		position[char] = index
	}
	// When the end of the string is reached check if the currentSubstringLength is longer than the 	longestSubstringLength
	if currentSubstringLength > longestSubstringLength {
		longestSubstringLength = currentSubstringLength
	}
	
	return longestSubstringLength
	
}
