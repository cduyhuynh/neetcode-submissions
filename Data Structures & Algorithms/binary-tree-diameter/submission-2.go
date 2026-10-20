/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}
	left_depth := height(root.Left)
	right_depth := height(root.Right)
	return max(left_depth + right_depth, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func height (head *TreeNode) int {
	if head == nil {
		return 0
	}
	return max(height(head.Left), height(head.Right)) + 1
}
