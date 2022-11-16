package main

func main() {
}

func getRC(id int, n int) (int, int) {
	r, c := (id-1)/n, (id-1)%n
	if r%2 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return r, c
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	book := make([]bool, n*n+1)
	type pair struct {
		id   int
		step int
	}

	q := make([]pair, 0)
	q = append(q, pair{
		id:   1,
		step: 0,
	})

	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		for i := 1; i <= 6; i++ {
			nextId := node.id + i
			if nextId > n*n {
				break
			}

			r, c := getRC(nextId, n)
			if board[r][c] != -1 {
				nextId = board[r][c]
			}

			if nextId == n*n {
				return node.step + 1
			}

			if !book[nextId] {
				book[nextId] = true
				q = append(q, pair{
					id:   nextId,
					step: node.step + 1,
				})
			}
		}
	}

	return -1
}
