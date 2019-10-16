package add_two_numbers

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 数组转化为ListNode
func arrayToListNode(a []int) *ListNode {
	node := &ListNode{}
	temp := &ListNode{}

	for index, value := range a {
		if index == 0 {
			temp.Val = value
			node = temp
		} else {
			temp.Next = &ListNode{Val: value}
			temp = temp.Next
		}
	}
	return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var isLeftAdd bool
	node := &ListNode{}
	tempNode := &ListNode{}
	for {
		temp := l1.Val + l2.Val
		if isLeftAdd {
			temp += 1
		}
		result := temp - 10
		var value int
		if result > 0 {
			isLeftAdd = true
			value = result
		} else {
			value = temp
		}
		tempNode.Val = value

	}

	return &ListNode{}
}
