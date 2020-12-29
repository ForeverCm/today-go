package main


func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func max(l, r int) int {
	if l <= r {
		return r
	}
	return l
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)), height(root.Left) + height(root.Right))
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}

