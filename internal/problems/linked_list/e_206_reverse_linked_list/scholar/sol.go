package scholar

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution() {
	n5 := &ListNode{Val: 5, Next: nil}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	reverseList(n1)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	next := head.Next
//
//	for {
//		prev = curr
//		curr = curr.Next
//
//		next = curr.Next
//		curr.Next = prev
//
//
//	}
//
//	return curr
//}
