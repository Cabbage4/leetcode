package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	q := []*Node{root}

	result := 1
	cur := 1
	next := 0
	for len(q) != 0 {
		cur--
		n := q[0]
		q = q[1:]

		if n.Children != nil {
			for _, v := range n.Children {
				q = append(q, v)
				next++
			}
		}

		if cur == 0 {
			if next == 0 {
				break
			}

			result++
			cur = next
			next = 0
		}
	}

	return result
}
