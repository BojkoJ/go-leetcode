package main

import "fmt"

// ------------------------------------------------------------------
// LeetCode: 1. Two Sums
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
// You can return the answer in any order.
// ------------------------------------------------------------------
// Solution #1 with time complexity O(n^2) - because two nested loops
//
//	func twoSum(nums []int, target int) []int {
//		for i, num := range nums {
//			for j, num2 := range nums {
//				if i <= j {
//					continue
//				}
//				if (num + num2) == target {
//					return []int{i, j}
//				}
//			}
//		}
//		return []int{}
//	}
//
// ------------------------------------------------------------------
// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
// ------------------------------------------------------------------
// Solution #2 with time complexity O(n) - because only one loop
// Trade-off: we make the time complexity better to ~O(n), but we pay with memory O(n), because map can have n items in the wors case.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // hash map will be saving: number -> index
	for i, num := range nums {
		// we calculate the complement - that's the number we need to add to "num", to get "target"
		complement := target - num
		// now we just look, if we have an item with key of "complement" in our map
		val, ok := m[complement] // we ask the Go hash map, if it already has an item with key "complement", val is then the value stored under that key (key->value) => (number->index)
		if ok {
			// it's in the map
			return []int{val, i}
		}
		// it's not in the map, save the index of that num to the map for next iterations and continue
		m[num] = i
	}
	return []int{}
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
