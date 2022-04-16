package main

import "fmt"

func main() {
	one := &Node{Val: 1, Neighbors: make([]*Node, 0)}
	two := &Node{Val: 2, Neighbors: make([]*Node, 0)}
	three := &Node{Val: 3, Neighbors: make([]*Node, 0)}
	four := &Node{Val: 4, Neighbors: make([]*Node, 0)}
	one.Neighbors = append(one.Neighbors, two, four)
	two.Neighbors = append(two.Neighbors, one, three)
	three.Neighbors = append(three.Neighbors, two, four)
	four.Neighbors = append(four.Neighbors, one, three)

	r := cloneGraph(one)
	fmt.Println(r)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return dfs(node, make(map[int]*Node))
}

func dfs(node *Node, mp map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, has := mp[node.Val]; has {
		return v
	}

	r := &Node{Val: node.Val}
	mp[node.Val] = r

	if node.Neighbors == nil {
		return r
	}

	r.Neighbors = make([]*Node, len(node.Neighbors))
	for i := 0; i < len(node.Neighbors); i++ {
		r.Neighbors[i] = dfs(node.Neighbors[i], mp)
	}
	return r
}
