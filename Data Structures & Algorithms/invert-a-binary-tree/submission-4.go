/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head == nil {
			continue
		}
		tmp := head.Left
		head.Left = head.Right
		head.Right = tmp

		queue = append(queue, head.Left, head.Right)
	}

    return root
}
