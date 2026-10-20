/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	longest := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left_depth := dfs(root.Left)
		right_depth := dfs(root.Right)
		longest = max(longest, left_depth + right_depth)
		return 1 + max(left_depth, right_depth)
	}
	dfs(root)
	return longest
}
