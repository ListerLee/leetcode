package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	ans := l3
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l3.Next = l2
			l2 = l2.Next
		} else {
			l3.Next = l1
			l1 = l1.Next
		}
		l3 = l3.Next
	}

	if l1 == nil {
		l3.Next = l2
	} else {
		l3.Next = l1
	}
	return ans.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
