package week03

type ListNode struct {
	Val  int
	Next *ListNode
}

var temp *ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	var merge func(int, int) *ListNode
	merge = func(left int, right int) *ListNode {
		if right == left {
			return lists[left]
		}
		mid := left + (right-left)/2
		listsA := merge(left, mid)
		listsB := merge(mid+1, right)
		return mergeTwoLists(listsA, listsB)
	}

	return merge(0, n-1)

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	last := &ListNode{}
	entrance := &ListNode{Next: last}
	for list1 != nil || list2 != nil {
		//什么时候要取l1的值
		if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
			last.Next = list1
			last = last.Next
			list1 = list1.Next
		} else {
			last.Next = list2
			last = last.Next
			list2 = list2.Next
		}
	}
	return entrance.Next.Next
}
