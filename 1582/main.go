package main

func main() {
}

func numSpecial(mat [][]int) int {
	r := 0
	for i := 0; i < len(mat); i++ {
		aj := -1
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				if aj != -1 {
					aj = -1
					break
				}
				aj = j
			}
		}

		if aj == -1 {
			continue
		}

		flag := false
		for ai := 0; ai < len(mat); ai++ {
			if ai == i {
				continue
			}

			if mat[ai][aj] == 1 {
				flag = true
				break
			}
		}

		if flag {
			continue
		}

		r++
	}

	return r
}
