package main

func main() {
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var r [][]int

	q := make([]*Node, 1)
	q[0] = root
	for len(q) != 0 {
		size := len(q)
		var tr []int
		for i := 0; i < size; i++ {
			for _, v := range q[i].Children {
				q = append(q, v)
			}
			tr = append(tr, q[i].Val)
		}
		r = append(r, tr)
		q = q[size:]
	}

	return r
}
