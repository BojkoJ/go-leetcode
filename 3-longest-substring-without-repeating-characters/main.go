package main

import "fmt"

// ------------------------------------------------------------------
// LeetCode: 3. Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest without duplicate characters.
// Constraints:
//
//    0 <= s.length <= 5 * 104
//    s consists of English letters, digits, symbols and spaces.
// ------------------------------------------------------------------
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
// ------------------------------------------------------------------
//Example 2:
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
// ------------------------------------------------------------------
//Example 3:
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
// ------------------------------------------------------------------

func lengthOfLongestSubstring(s string) int {
	// This is the sliding window pattern used to track a continuous subset of elements.
	// We use it here because it allows us to find the longest substring in O(n) time by avoiding repetitive checks of all possible substrings.
	// Brute force would be easier to implement here BUT slower.
	// --------
	// Method of sliding window: there is imaginary "window" spreading from the left to the right of given string
	// During the reading we constatly update simple integer variable maxLen our currently found maximum
	// We use hash map not for saving the whole word, but as a quick map for characters:
	//     - have we seen this given character in this current window? On which index?
	//     - as key we use that given character, and it's value would be its index
	// When we find character which is already in the map, simply cut the beggining of the window:
	//     - that means move the begining of the window to the place right begind first occurence of this duplicit character

	maxLen, currentLen := 0, 0 // biggest and current lenght found of substríng of unique characters

	m := make(map[rune]int) // hash map will be saving: character -> index

	// start of our window
	start := 0

	// no need to create the "end" of our window - that will be the index in the loop

	// this will iterate through the characters of given string
	for index, character := range s {
		val, ok := m[character] // this pulls the val out of the map if the character key exists

		// Is the character in the map AND is its old position in our current window?
		if ok && val >= start {
			// if the character is in the map and its old position reaches inside the current window it means:
			// DUPLICITE - so we need to move the start of our window behind the current character
			// this moves the window and creates opportunity to check new substrings

			// move the start of the window
			start = val + 1 // one character to the right of current
		}

		// we need to save the character to the map every iteration
		// if it's already in the map, this will just update its index
		m[character] = index

		currentLen = index - start + 1 // lenght of current window is it's end minus it's start

		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

func main() {
	s := "pwwkew"
	res := lengthOfLongestSubstring(s)

	fmt.Printf("%d", res)
}
