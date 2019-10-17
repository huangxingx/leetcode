package invertTree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		tempNode := root.Left
		root.Left = root.Right
		root.Right = tempNode

		invertTree(root.Left)
		invertTree(root.Right)
	} else {
		return nil
	}

	return root
}
