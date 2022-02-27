package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(averageOfLevels(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queen []*TreeNode

func (s *Queen) Len() int {
	return len(*s)
}

func (s *Queen) Pop() *TreeNode {
	p := (*s)[0]
	*s = (*s)[1:]
	return p
}

func (s *Queen) Push(v *TreeNode) {
	*s = append(*s, v)
}

func (s *Queen) Peek() *TreeNode {
	return (*s)[s.Len()-1]
}

func averageOfLevels(root *TreeNode) []float64 {
	stack := new(Queen)
	stack.Push(root)

	cur := 1
	nextCur := 0

	result := make([]float64, 0)
	count := 1
	total := 0
	for stack.Len() != 0 {
		p := stack.Pop()
		cur--
		total += p.Val

		if p.Left != nil {
			nextCur++
			stack.Push(p.Left)
		}
		if p.Right != nil {
			nextCur++
			stack.Push(p.Right)
		}

		if cur == 0 {
			r, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(total)/float64(count)), 64)
			result = append(result, r)

			cur = nextCur
			count = cur
			total = 0
			nextCur = 0
		}
	}

	return result
}
