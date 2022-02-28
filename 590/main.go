package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	dfs(root, &result)
	return result
}

func dfs(root *Node, list *[]int) {
	if root.Children != nil {
		for _, v := range root.Children {
			dfs(v, list)
		}
	}

	*list = append(*list, root.Val)
}
