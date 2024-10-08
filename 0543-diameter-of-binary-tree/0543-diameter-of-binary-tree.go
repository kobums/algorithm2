/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)
	return res
}

func helper(root *TreeNode) (length, diameter int) {
	if root == nil {
		return
	}

	leftLen, leftDia := helper(root.Left)
	rightLen, rightDia := helper(root.Right)

	length = max(leftLen, rightLen) + 1
	diameter = max(leftLen+rightLen, max(leftDia, rightDia))
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}