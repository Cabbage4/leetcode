package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	r := make([][]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		tr := make([]int, 0)

		for i := 0; i < size; i++ {
			t := q[i]
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
			tr = append(tr, t.Val)
		}

		r = append([][]int{tr}, r...)
		q = q[size:]
	}

	return r
}
