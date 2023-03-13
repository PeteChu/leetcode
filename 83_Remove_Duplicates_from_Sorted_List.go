// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	currentNode := head

	for true {
		if currentNode == nil {
			break
		}

		if currentNode.Next != nil {
			currentVal := currentNode.Val
			var nextNode *ListNode
			nextNode = currentNode.Next
			for true {
				if currentVal != nextNode.Val {
					currentNode.Next = nextNode
					break
				}
				nextNode = nextNode.Next
				if nextNode == nil {
					currentNode.Next = nil
					break
				}
			}
		}
		currentNode = currentNode.Next
	}

	return head
}

func generateNodeList(vals []int) *ListNode {

	head := &ListNode{
		Val: vals[0],
	}
	headAddr := head

	for i := 1; i < len(vals); i++ {
		newNode := &ListNode{}
		head.Next = &ListNode{
			Val:  vals[i],
			Next: newNode,
		}
		head = head.Next
	}

	return headAddr
}

func printNodeList(head *ListNode) {
	for true {
		fmt.Println(head.Val)
		head = head.Next
		if head == nil {
			break
		}
	}
}

func main() {
	listNode := generateNodeList([]int{1, 1, 2, 2, 3})
	ans := deleteDuplicates(listNode)
	printNodeList(ans)
}
