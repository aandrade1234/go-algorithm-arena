package main

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	nodesToDelete := make(map[int]struct{}, len(to_delete))
	for _, val := range to_delete {
		nodesToDelete[val] = struct{}{}
	}

	var result []*TreeNode

	var processNode func(*TreeNode, bool) *TreeNode
	processNode = func(node *TreeNode, isRoot bool) *TreeNode {
		if node == nil {
			return nil
		}

		_, shouldDelete := nodesToDelete[node.Value]
		if isRoot && !shouldDelete {
			result = append(result, node)
		}

		node.LeftNode = processNode(node.LeftNode, shouldDelete)
		node.RightNode = processNode(node.RightNode, shouldDelete)

		if shouldDelete {
			return nil
		}

		return node
	}

	processNode(root, true)

	return result
}
