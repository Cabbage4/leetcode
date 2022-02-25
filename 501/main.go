package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	mp := make(map[int]int)
	dfs(root, mp)

	count := -1
	result := make([]int, 0)
	for v, c := range mp {
		if count == -1 || count == c {
			result = append(result, v)
			count = c
			continue
		}

		if count < c {
			result = []int{v}
			count = c
		}
	}

	return result
}

func dfs(root *TreeNode, mp map[int]int) {
	if root == nil {
		return
	}

	mp[root.Val]++
	dfs(root.Left, mp)
	dfs(root.Right, mp)
}
