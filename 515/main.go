package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := make([]*TreeNode, 1)
	q[0] = root

	var r []int
	for len(q) != 0 {
		size := len(q)

		maxValue := q[0].Val
		for i := 0; i < size; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}

			if q[i].Val > maxValue {
				maxValue = q[i].Val
			}
		}

		r = append(r, maxValue)
		q = q[size:]
	}

	return r
}
