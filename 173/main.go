package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	data  []int
	index int
}

func Constructor(root *TreeNode) BSTIterator {
	r := make([]int, 0)
	dfs(root, &r)
	return BSTIterator{
		data:  r,
		index: 0,
	}
}

func (b *BSTIterator) Next() int {
	r := b.data[b.index]
	b.index++
	return r
}

func (b *BSTIterator) HasNext() bool {
	return b.index < len(b.data)
}

func dfs(root *TreeNode, r *[]int) {
	if root.Left != nil {
		dfs(root.Left, r)
	}
	*r = append(*r, root.Val)
	if root.Right != nil {
		dfs(root.Right, r)
	}
}
