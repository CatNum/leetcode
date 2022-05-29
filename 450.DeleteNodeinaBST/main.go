package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	// root为要删除的节点
	if root.Val == key {
		// 1. 左孩子为空，返回右孩子
		if root.Left == nil {
			return root.Right
		}
		// 2. 右孩子为空，返回左孩子
		if root.Right == nil {
			return root.Left 
		}
		// 3. 左右孩子不为空，找到右孩子的最左侧孩子
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		// 把root左孩子，接到最左侧孩子的左侧
		cur.Left = root.Left
		// 把root更新为root的右孩子
		root = root.Right
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}

	return root
}