package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	q := make([]*TreeNode, 1)
	q[0] = root

	var r int
	for len(q) != 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		if len(q) == size {
			r = q[0].Val
			break
		}

		q = q[size:]
	}

	return r
}
