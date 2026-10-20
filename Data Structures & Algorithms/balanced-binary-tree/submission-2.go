/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	balanced := true
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := 1 + dfs(node.Left)
		right := 1 + dfs(node.Right)
		if (left - right >= -1) && (left - right <= 1) {
			return max(left, right)
		} else {
			balanced = false
			return 0
		}
	}
	dfs(root)
	return balanced
}