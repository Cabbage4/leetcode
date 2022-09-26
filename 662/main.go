package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := make([]Pair, 1)
	q[0] = Pair{
		node:  root,
		index: 1,
	}

	r := 1
	for len(q) != 0 {
		tmpR := q[len(q)-1].index - q[0].index + 1
		if tmpR > r {
			r = tmpR
		}

		size := len(q)

		for i := 0; i < size; i++ {
			v := q[i]
			if v.node.Left != nil {
				q = append(q, Pair{
					node:  v.node.Left,
					index: 2 * v.index,
				})
			}
			if v.node.Right != nil {
				q = append(q, Pair{
					node:  v.node.Right,
					index: 2*v.index + 1,
				})
			}
		}

		q = q[size:]
	}

	return r
}
