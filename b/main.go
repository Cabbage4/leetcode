package main

import (
	"fmt"
)

func main() {
	creators := []string{"alice"}
	ids := []string{"a"}
	views := []int{0}
	fmt.Println(mostPopularCreator(creators, ids, views))
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	var max int
	var maxList []int
	imp := make(map[string][]int)
	cmp := make(map[string]int)
	for i := range creators {
		imp[creators[i]] = append(imp[creators[i]], i)
		cmp[creators[i]] += views[i]
		if cmp[creators[i]] > max {
			max = cmp[creators[i]]
			maxList = []int{i}
		} else if cmp[creators[i]] == max {
			maxList = append(maxList, i)
		}
	}

	var r [][]string
	for _, i := range maxList {
		list := imp[creators[i]]
		maxCount := 0
		maxValue := ""
		for j := range list {
			if views[list[j]] > maxCount {
				maxCount = views[list[j]]
				maxValue = ids[list[j]]
			} else if views[list[j]] == maxCount && (ids[list[j]] < maxValue || maxValue == "") {
				maxValue = ids[list[j]]
			}
		}

		r = append(r, []string{creators[i], maxValue})
	}
	return r
}
