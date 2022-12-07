package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(byte(bits.TrailingZeros(0b000010000)))
}

func solveSudoku(board [][]byte) {
	var lines, columns [9]int
	var blocks [3][3]int

	flip := func(i, j int, d byte) {
		lines[i] ^= 1 << d
		columns[j] ^= 1 << d
		blocks[i/3][j/3] ^= 1 << d
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				flip(i, j, board[i][j]-'1')
			}
		}
	}

	for {
		modified := false

		for i := range board {
			for j := range board[i] {
				if board[i][j] != '.' {
					continue
				}

				mask := 0x1ff & (^uint(lines[i] | columns[j] | blocks[i/3][j/3]))
				if mask&(mask-1) == 0 {
					d := byte(bits.TrailingZeros(mask))
					board[i][j] = d + '1'
					flip(i, j, d)
					modified = true
				}
			}
		}

		if !modified {
			break
		}
	}

	q := make([][2]int, 0)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				q = append(q, [2]int{i, j})
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(q) {
			return true
		}

		i, j := q[pos][0], q[pos][1]
		mask := 0x1ff & (^uint(lines[i] | columns[j] | blocks[i/3][j/3]))
		for ; mask > 0; mask &= mask - 1 {
			d := byte(bits.TrailingZeros(mask))
			board[i][j] = d + '1'
			flip(i, j, d)
			if dfs(pos + 1) {
				return true
			}
			flip(i, j, d)
		}

		return false
	}

	dfs(0)
}
