package vscode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 数组转链表
func makeListNode(elements []int) *ListNode {
	if len(elements) == 0 {
		return nil
	}
	res := &ListNode{
		Val: elements[0],
	}
	// temp对res复制，二者没有关系了
	temp := res
	for i := 1; i < len(elements); i++ {
		// 对temp的子属性赋值
		temp.Next = &ListNode{Val: elements[i]}
		// 再将temp指向当前子属性 作为下一轮循环的当前值
		temp = temp.Next
	}

	return res
}

func changeListNode2List(listNode *ListNode) []int {
	ret := make([]int, 0)
	ret = append(ret, listNode.Val)
	for listNode.Next != nil {
		listNode = listNode.Next
		ret = append(ret, listNode.Val)
	}
	return ret
}
