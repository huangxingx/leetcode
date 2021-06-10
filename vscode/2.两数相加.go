package vscode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	tmp := root
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			l1 = new(ListNode)
		}
		if l2 == nil {
			l2 = new(ListNode)
		}
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		sumNode := &ListNode{
			Val: sum % 10,
		}
		tmp.Next = sumNode
		tmp = sumNode

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return root.Next
}

// @lc code=end
