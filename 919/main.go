package main

import "fmt"

func main() {
	c := Constructor(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}})
	fmt.Println(c.Insert(3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
	q    []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	if root == nil {
		return CBTInserter{
			root: nil,
			q:    nil,
		}
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	for {
		if q[0].Left == nil || q[0].Right == nil {
			break
		}
		q = append(q, q[0].Left, q[0].Right)
		q = q[1:]
	}

	return CBTInserter{
		root: root,
		q:    q,
	}
}

func (c *CBTInserter) Insert(val int) int {
	if c.q[0].Left == nil {
		c.q[0].Left = &TreeNode{
			Val: val,
		}
		return c.q[0].Val
	}

	if c.q[0].Right == nil {
		c.q[0].Right = &TreeNode{
			Val: val,
		}

		r := c.q[0].Val

		c.q = append(c.q, c.q[0].Left, c.q[0].Right)
		c.q = c.q[1:]

		return r
	}

	return -1
}

func (c *CBTInserter) Get_root() *TreeNode {
	return c.root
}
