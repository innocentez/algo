package scholar

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution() {
	a5 := &ListNode{Val: 5, Next: nil}
	a4 := &ListNode{Val: 4, Next: a5}
	a3 := &ListNode{Val: 3, Next: a4}
	a2 := &ListNode{Val: 2, Next: a3}
	a1 := &ListNode{Val: 1, Next: a2}

	b5 := &ListNode{Val: 5, Next: nil}
	b4 := &ListNode{Val: 4, Next: b5}
	b3 := &ListNode{Val: 3, Next: b4}
	b2 := &ListNode{Val: 2, Next: b3}
	b1 := &ListNode{Val: 1, Next: b2}

	mergeTwoLists(a1, b1)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	for {

	}
}
