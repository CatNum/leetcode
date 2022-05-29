package main
//二叉树遍历
import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//前序遍历递归实现
func dpreorderTraversal(root *TreeNode) (res []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode){
		if node == nil {
			return
		}
		res = append(res,node.Val)
		//左子树
		preorder(node.Left)
		//右子树
		preorder(node.Right)
	}
	preorder(root)
	return
}
//中序遍历递归实现
func dinorderTraversal(root *TreeNode) (res []int) {
	//为什么不能直接用inorder := func(node *TreeNode){}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		//如果有左子树结点，就进入
		inorder(node.Left)
		res = append(res, node.Val)
		//如果有右子树结点，就进入
		inorder(node.Right)
	}
	inorder(root)
	return
}
//后序遍历递归实现
func dpostorderTraversal(root *TreeNode) (res []int) {
	var dpost func(node *TreeNode)
	dpost = func(node *TreeNode){
		if node == nil {
			return
		}
		//左子树
		dpost(node.Left)
		//右子树
		dpost(node.Right)
		res = append(res,node.Val)
	}
	dpost(root)
	return
}
//后序遍历迭代实现
func postorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

func main() {
	var t0,t1,t2,t3,t4,t5,t6 TreeNode
	t0.Val = 0
	t1.Val = 1
	t2.Val = 2
	t3.Val = 3
	t4.Val = 4
	t5.Val = 5
	t6.Val = 6
	t0.Left = &t1
	t0.Right = &t2
	t1.Left = &t3
	t1.Right = &t4
	t2.Left = &t5
	t2.Right = &t6
	//递归前序遍历
	dpr:=dpreorderTraversal(&t0)
	fmt.Println("递归前序遍历：",dpr)
	//递归中序遍历
	din := dinorderTraversal(&t0)
	fmt.Println("递归中序遍历",din)
	//递归后序遍历
	dpo := dpostorderTraversal(&t0)
	fmt.Println("递归后序遍历",dpo)
}
