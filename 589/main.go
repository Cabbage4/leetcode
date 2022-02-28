package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	dfs(root, &result)
	return result
}

func dfs(root *Node, list *[]int) {
	*list = append(*list, root.Val)
	if root.Children != nil {
		for _, v := range root.Children {
			dfs(v, list)
		}
	}
}
