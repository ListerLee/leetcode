package removeNthNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getNodeLength(head *ListNode) (cnt int) {
	cnt = 0
	cur := head
	for ; cur != nil; cur = cur.Next {
		cnt++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	length := getNodeLength(head)
	dummy := &ListNode{0, head}
	cur := dummy

	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return dummy.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
