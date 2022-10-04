package main

func main() {
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(tree1 *Node, tree2 *Node) *Node {
	if tree1.IsLeaf && tree1.Val {
		return tree1
	}

	if tree2.IsLeaf && tree2.Val {
		return tree2
	}

	if tree1.IsLeaf && !tree1.Val {
		return tree2
	}

	if tree2.IsLeaf && !tree2.Val {
		return tree1
	}

	r := &Node{
		TopLeft:     intersect(tree1.TopLeft, tree2.TopLeft),
		TopRight:    intersect(tree1.TopRight, tree2.TopRight),
		BottomLeft:  intersect(tree1.BottomLeft, tree2.BottomLeft),
		BottomRight: intersect(tree1.BottomRight, tree2.BottomRight),
	}

	if r.TopLeft.IsLeaf &&
		r.TopRight.IsLeaf &&
		r.BottomLeft.IsLeaf &&
		r.BottomRight.IsLeaf &&
		r.TopLeft.Val == r.TopRight.Val &&
		r.TopLeft.Val == r.BottomLeft.Val &&
		r.TopLeft.Val == r.BottomRight.Val {
		return &Node{
			Val:    r.TopLeft.Val,
			IsLeaf: true,
		}
	}

	return r
}
