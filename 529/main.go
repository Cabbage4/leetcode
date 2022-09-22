package main

import "fmt"

func main() {
	testData := [][]byte{{'E'}}
	testData = updateBoard(testData, []int{0, 0})
	for i := 0; i < len(testData); i++ {
		for j := 0; j < len(testData[i]); j++ {
			fmt.Printf("%c ", testData[i][j])
		}
		fmt.Println()
	}
}

func updateBoard(board [][]byte, click []int) [][]byte {
	directs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	book := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		book[i] = make([]bool, len(board[i]))
	}

	var dfs func(int, int)
	dfs = func(x int, y int) {
		c := 0
		for _, d := range directs {
			nx, ny := x+d[0], y+d[1]
			if !(0 <= nx && nx < len(board) && 0 <= ny && ny < len(board[0])) {
				continue
			}
			if board[nx][ny] == 'M' {
				c++
				continue
			}
		}

		if c != 0 {
			board[x][y] = byte(c + '0')
			return
		}

		for _, d := range directs {
			nx, ny := x+d[0], y+d[1]
			if !(0 <= nx && nx < len(board) && 0 <= ny && ny < len(board[0])) {
				continue
			}

			if book[nx][ny] {
				continue
			}

			if board[nx][ny] == 'E' {
				book[nx][ny] = true
				board[nx][ny] = 'B'
				dfs(nx, ny)

				continue
			}
		}
	}

	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	} else {
		board[click[0]][click[1]] = 'B'
	}

	dfs(click[0], click[1])

	return board
}
