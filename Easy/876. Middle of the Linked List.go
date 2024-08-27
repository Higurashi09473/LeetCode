package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var count int
	g := head
	for g != nil {
		g = g.Next
		count++
	}
	for i := 0; i < count/2; i++ {
		head = head.Next
	}
	return head
}
