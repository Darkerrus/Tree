package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := findCommonAncestor(root.Left, p, q)
	right := findCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		
	}

	nodeA := root.Left.Left.Left.Left.Left  // Узел A со значением 7
	nodeB := root.Left.Left.Right // Узел B со значением 5

	commonAncestor := findCommonAncestor(root, nodeA, nodeB)

	if commonAncestor != nil {
		fmt.Printf("Общий узел С со значением %d\n", commonAncestor.Val)
	} else {
		fmt.Println("Общий узел С не найден")
	}
}
