/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q == nil) {
		return true
	} else if (p != nil && q != nil) {
	} else {
		return false
	}
    p_queue := []*TreeNode{p}
	q_queue := []*TreeNode{q}

	for len(p_queue) > 0 && len(q_queue) > 0 {
		p_node := p_queue[0]
		q_node := q_queue[0]
		if p_node.Val != q_node.Val {
			return false
		}

		p_queue = p_queue[1:]
		q_queue = q_queue[1:]
		if p_node.Left != nil && q_node.Left != nil{
			p_queue = append(p_queue, p_node.Left)
			q_queue = append(q_queue, q_node.Left)
		} else if p_node.Left == nil && q_node.Left == nil{
		} else {
			return false
		}
		if p_node.Right != nil && q_node.Right != nil {
			p_queue = append(p_queue, p_node.Right)
			q_queue = append(q_queue, q_node.Right)
		} else if p_node.Right == nil && q_node.Right == nil{
		} else {
			return false
		}
	}

	return true
}
