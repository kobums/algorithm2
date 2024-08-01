/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxVal int) int {
    if node == nil {
        return 0
    }

    ans := 0
    if node.Val >= maxVal {
        ans = 1
    }
    maxVal = max(maxVal, node.Val)
    ans += dfs(node.Left, maxVal)
    ans += dfs(node.Right, maxVal)

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
