package main

import "fmt"

func main() {
	//data := [][]int{{0, 1}, {1, 0}}
	data := [][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}}
	root := construct(data)
	fmt.Println(root)
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	preSum := make([][]int, len(grid)+1)
	preSum[0] = make([]int, len(grid[0])+1)
	for i := range grid {
		preSum[i+1] = make([]int, len(grid[0])+1)
		for j := range grid[i] {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + grid[i][j]
		}
	}

	var dfs func(left, right, up, down int) *Node
	dfs = func(left, right, up, down int) *Node {
		total := preSum[down+1][right+1] - preSum[up][right+1] - preSum[down+1][left] + preSum[up][left]
		if total == 0 || total == (right-left+1)*(down-up+1) {
			return &Node{
				Val:    grid[up][left] == 1,
				IsLeaf: true,
			}
		}

		if right == left {
			return &Node{
				Val:    grid[up][left] == 1,
				IsLeaf: true,
			}
		}

		root := new(Node)

		lrMid := (left + right) / 2
		udMid := (up + down) / 2

		root.TopLeft = dfs(left, lrMid, up, udMid)
		root.TopRight = dfs(lrMid+1, right, up, udMid)
		root.BottomLeft = dfs(left, lrMid, udMid+1, down)
		root.BottomRight = dfs(lrMid+1, right, udMid+1, down)

		return root
	}

	return dfs(0, len(grid[0])-1, 0, len(grid)-1)
}
