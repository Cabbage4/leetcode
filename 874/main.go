package main

import "fmt"

func main() {
	fmt.Println(robotSim([]int{-2, -1, -2, 3, 7}, [][]int{{1, -3}, {2, -3}, {4, 0}, {-2, 5}, {-5, 2}, {0, 0}, {4, -4}, {-2, -5}, {-1, -2}, {0, 2}}))
	fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}))
	fmt.Println(robotSim([]int{6, -1, -1, 6}, [][]int{}))
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}

func robotSim(commands []int, obstacles [][]int) int {
	directs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	curForward := 0
	r := 0
	x, y := 0, 0

	for _, c := range commands {
		if c == -1 {
			curForward = (curForward + 1) % 4
			continue
		} else if c == -2 {
			curForward = (4 + curForward - 1) % 4
			continue
		}

		direct := directs[curForward]
		nx, ny := x+direct[0]*c, y+direct[1]*c

		for _, o := range obstacles {
			ox, oy := o[0], o[1]
			if ox == 0 && oy == 0 && x == 0 && y == 0 {
				continue
			}

			xLeft, xRight := min(nx, x), max(nx, x)
			yBottom, yTop := min(ny, y), max(ny, y)
			if xLeft <= ox && ox <= xRight && yBottom <= oy && oy <= yTop {
				switch curForward {
				case 0:
					ny = oy - 1
				case 1:
					nx = ox - 1
				case 2:
					ny = oy + 1
				case 3:
					nx = ox + 1
				}
			}
		}

		x, y = nx, ny

		if tmp := x*x + y*y; tmp > r {
			r = tmp
		}
	}

	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
