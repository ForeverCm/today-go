package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1))
	//prints: dir1 => AAAA fmt.Println("dir2 =>",string(dir2))
	//prints: dir2 => BBBBBBBBB dir1 = append(dir1,"suffix"...)
	//fmt.Println("dir1 =>",string(dir1))
	//prints: dir1 => AAAAsuffix fmt.Println("dir2 =>",string(dir2))
	//prints: dir2 => uffixBBBB
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

