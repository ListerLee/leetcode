package mergeK

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tail.Next = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			l1 = l1.Next
		}

		tail = tail.Next
	}

	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}

	return head.Next
}

func mergeKLists1(lists []*ListNode) *ListNode {
	var ans *ListNode = nil
	for _, list := range lists {
		ans = mergeTwoLists(ans, list)
	}

	return ans
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2

	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}
