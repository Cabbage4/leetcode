package main

import (
	"fmt"
)

func main() {
	fmt.Println(predictPartyVictory("DRRDRDRDRDDRDRDR"))
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
}

func predictPartyVictory(senate string) string {
	ri := make([]int, 0)
	di := make([]int, 0)
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			ri = append(ri, i)
		} else {
			di = append(di, i)
		}
	}

	for len(ri) > 0 && len(di) > 0 {
		if ri[0] < di[0] {
			ri = append(ri, ri[0]+len(senate))
		} else {
			di = append(di, di[0]+len(senate))
		}
		di = di[1:]
		ri = ri[1:]
	}

	if len(ri) == 0 {
		return "Dire"
	}

	return "Radiant"
}
