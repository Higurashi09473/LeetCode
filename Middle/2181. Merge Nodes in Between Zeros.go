package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	var count int
	l := &ListNode{}
	t := l
	for head != nil {
		if count != 0 && head.Val == 0 {
			t.Next = &ListNode{Val: count}
			t = t.Next
			count = 0
		}
		count += head.Val
		head = head.Next
	}
	return l.Next
}
