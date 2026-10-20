/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	if isBalanced(root.Left) && isBalanced(root.Right) {
		lheight := height(root.Left)
		rheight := height(root.Right)
		if (lheight <= rheight + 1 && lheight >= rheight - 1) || (rheight <= lheight + 1 && rheight >= lheight - 1) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lheight := height(root.Left)
	rheight := height(root.Right)
	return 1 + max(lheight, rheight)
}