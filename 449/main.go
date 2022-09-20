package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	list := make([]string, 0)
	postOrder(root, &list)

	return strings.Join(list, " ")
}

func postOrder(root *TreeNode, list *[]string) {
	if root == nil {
		return
	}

	postOrder(root.Left, list)
	postOrder(root.Right, list)

	*list = append(*list, strconv.Itoa(root.Val))
}

func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	list := strings.Split(data, " ")
	return rebuild(&list, math.MinInt32, math.MaxInt32)
}

func rebuild(list *[]string, min, max int) *TreeNode {
	if len(*list) == 0 {
		return nil
	}

	value, _ := strconv.Atoi((*list)[len(*list)-1])

	if value < min || value > max {
		return nil
	}
	
	*list = (*list)[:len(*list)-1]
	return &TreeNode{
		Val:   value,
		Right: rebuild(list, value, max),
		Left:  rebuild(list, min, value),
	}
}
