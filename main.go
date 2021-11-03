package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res

	l1 = &ListNode{Next: l1}
	l2 = &ListNode{Next: l2}

	for *l1.Next != (ListNode{}) {
		l1 = l1.Next
		l2 = l2.Next

		sum := curr.Val + l1.Val + l2.Val
		mod := sum % 10

		curr.Val = mod
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = (sum - mod) / 10
	}

	return res
}

func main() {
	res := addTwoNumbers(
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{}}}},
		&ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{}}}},
	)

	fmt.Println(res.Val, res.Next.Val, res.Next.Next.Val)
}
