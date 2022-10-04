package main

import "fmt"

func main() {
	data, target := []string{"0201", "0101", "0102", "1212", "2002"}, "0202"
	//data, target := []string{"8888"}, "0009"
	fmt.Println(openLock(data, target))
}

type pair struct {
	status string
	step   int
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	mp := make(map[string]bool)
	for _, d := range deadends {
		mp[d] = true
	}

	if mp["0000"] {
		return -1
	}

	book := make(map[string]bool)
	q := make([]*pair, 1)
	q[0] = &pair{
		status: "0000",
		step:   0,
	}

	getStatus := func(str string) []string {
		var r []string
		strByte := []byte(str)
		for i := 0; i < len(strByte); i++ {
			tmp := strByte[i] - '0'
			strByte[i] = (10+tmp-1)%10 + '0'
			r = append(r, string(strByte))
			strByte[i] = (10+tmp+1)%10 + '0'
			r = append(r, string(strByte))
			strByte[i] = tmp + '0'
		}
		return r
	}

	book["0000"] = true

	for len(q) != 0 {
		cur := q[0]

		for _, p := range getStatus(cur.status) {
			if book[p] || mp[p] {
				continue
			}

			if p == target {
				return cur.step + 1
			}

			book[p] = true
			q = append(q, &pair{
				status: p,
				step:   cur.step + 1,
			})
		}

		q = q[1:]
	}

	return -1
}
