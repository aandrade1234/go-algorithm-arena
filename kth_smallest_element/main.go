package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}

	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(kthSmallest1(root1, 1))
	fmt.Println(kthSmallest1(root2, 3))
	fmt.Println()
	fmt.Println(kthSmallest2(root1, 1))
	fmt.Println(kthSmallest2(root2, 3))
}

func kthSmallest1(root *TreeNode, k int) int {
	result := -1
	count := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		count++
		if count == k {
			result = node.Val

			return
		}

		inorder(node.Right)
	}

	inorder(root)

	return result
}

func kthSmallest2(root *TreeNode, k int) int {
	nums := DFS(root, nil)

	return nums[k-1]
}

func DFS(node *TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}

	arr = DFS(node.Left, arr)
	arr = append(arr, node.Val)
	arr = DFS(node.Right, arr)

	return arr
}
