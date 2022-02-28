package main

import "fmt"

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(floodFill(image, 1, 1, 2))
}

var direct = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	queue := make([][]int, 0)
	queue = append(queue, []int{sr, sc})

	book := make([][]int, len(image))
	for i := range image {
		book[i] = make([]int, len(image[0]))
	}

	for len(queue) != 0 {
		node := queue[0]
		x := node[0]
		y := node[1]
		queue = queue[1:]

		for _, v := range direct {
			nx := x + v[0]
			ny := y + v[1]

			if nx >= 0 && nx < len(image) &&
				ny >= 0 && ny < len(image[0]) &&
				book[nx][ny] == 0 &&
				image[nx][ny] == image[x][y] {
				book[nx][ny] = 1
				queue = append(queue, []int{nx, ny})
			}
		}
		image[x][y] = newColor
		book[x][y] = 2
	}

	return image
}
