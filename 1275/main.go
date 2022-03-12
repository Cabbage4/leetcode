package main

func main() {
}

func tictactoe(moves [][]int) string {
	g := make([][]int, 3)
	for i := 0; i < 3; i++ {
		g[i] = make([]int, 3)
	}

	turn := 1
	for _, v := range moves {
		g[v[0]][v[1]] = turn
		if turn == 1 {
			turn = 2
		} else if turn == 2 {
			turn = 1
		}
	}

	list := [][]int{
		{g[0][0], g[1][1], g[2][2]},
		{g[0][2], g[1][1], g[2][0]},
	}
	for i := 0; i < 3; i++ {
		list = append(list, g[i])
		list = append(list, []int{g[0][i], g[1][i], g[2][i]})
	}

	for _, v := range list {
		if r := check(v); r != -1 {
			if r == 1 {
				return "A"
			} else if r == 2 {
				return "B"
			}
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func check(list []int) int {
	if list[0] == list[1] && list[1] == list[2] {
		return list[1]
	}

	return -1
}
