package week01

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	entrance := &ListNode{Next: newList}
	for l1 != nil || l2 != nil {
		//什么时候入l1 :  l2出界 或者  l1和l2都没有出界，l1<=l2
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			newList.Next = l1
			newList = l1
			l1 = l1.Next
		} else {
			newList.Next = l2
			newList = l2
			l2 = l2.Next
		}
	}
	return entrance.Next.Next
}
