/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		head := queue[0]
		fmt.Println("head", head.Val)
		tmp := head.Left
		head.Left = head.Right
		head.Right = tmp

		queue = queue[1:]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}

    return root
}
