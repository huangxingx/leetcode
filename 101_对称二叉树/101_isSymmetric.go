package isSymmetric

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return _isSymmetric(root, root)
}

func _isSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil || right != nil {
		return false
	}

	return _isSymmetric(left.Left, right.Right) && _isSymmetric(right.Left, left.Right) && left.Val == right.Val

}
