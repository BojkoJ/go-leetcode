package main

import "fmt"

// ------------------------------------------------------------------
// LeetCode: 2. Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Constraints:
//    The number of nodes in each linked list is in the range [1, 100].
//    0 <= Node.val <= 9
//    It is guaranteed that the list represents a number that does not have leading zeros.
// ------------------------------------------------------------------
// Example 1:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// ------------------------------------------------------------------
// Example 2:
// Input: l1 = [0], l2 = [0]
// Output: [0]
// ------------------------------------------------------------------
// Example 3:
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
// ------------------------------------------------------------------

type ListNode struct {
	Val  int
	Next *ListNode
	// Why Pointer here? Because we want to modify the Next of the current node,
	// and if we use a non-pointer type, we would be working with a copy of the node,
	// which would not affect the original linked list.
	// By using a pointer, we can directly modify the Next of the current node in the original linked list.
}

// To the function we pass only first Node of the first LinkedList as l1 and first node of the second one
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// One of the constraints says that the number of nodes can be even 100!
	// If we would just glue them to a variable, it would be a number with 100 digits,
	// which is not possible to store in a standard integer type in Go. - it's called integer overflow.

	// create dummy head
	head := &ListNode{Val: 0}
	// current is the pointer which we will use to move in the new, result list
	current := head
	carry := 0 // This variable will store the carry from the previous addition, it starts with 0 because there is no carry at the beginning.

	// l1 is a pointer to the first node of the first linked list, l2 is a pointer to the first node of the second linked list.
	// what we will do is move the pointer to its child until there are child nodes available

	// this means: we will loop while l1 is not nil or l2 is not nil or carry is not 0
	// so if l1 is nil, but l2 is not nil (different lengths of linked lists) or if there is a carry from the previous addition, we will continue the loop

	// the loop will end only when both linked list pointers are nil and there is no carry left to add
	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2 + carry
		carry = sum / 10  // 13 / 10 = 1
		digit := sum % 10 // we keep only the last digit of the sum, because we can only store one digit in each node: 13 % 10 = 3 => we kept only the last digit
		current.Next = &ListNode{Val: digit}
		// move our new list pointer
		current = current.Next // this says: new current will point to the next node of the old current
		// now we move the pointers of both the lists
		if l1 != nil {
			// this means the previous move of the pointer didn't move it to nil child node, so we move it again until we move it to empy child node
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// now just return pointer to everything after dummy head - so dummy head's child
	return head.Next
}

func main() {
	// Let's test example 1:
	// l1 = [2,4,3], l2 = [5,6,4]

	// Function expects pointers, thus we need to create these variables as pointers, we use auto-assign
	// Why function expects pointers: There is a golden rule that when we pass struct to a function in Go, in a majority of cases
	// the preferable way is to do it with pointer, so the function does not create a copy of it and so it's more memory-efficient.
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	// Why not just say: l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}?
	// and then just pass l1's address with &l1 to the function?

	// Again: result is a pointer, not a value, because we want to return a pointer to the head of the resulting linked list.
	result := addTwoNumbers(l1, l2)

	// now let's print the result in the same manner

	fmt.Print("[")
	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
		if result != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}
