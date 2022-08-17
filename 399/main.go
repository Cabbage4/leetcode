package main

import "fmt"

func main() {
	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}},
		[]float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))

	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
		[]float64{1.5, 2.5, 5.0},
		[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
}

type Item struct {
	Value float64
	Key   string
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	mp := make(map[string][]Item)
	set := make(map[string]bool)
	for i, equation := range equations {
		set[equation[0]] = true
		set[equation[1]] = true

		mp[equation[0]] = append(mp[equation[0]], Item{
			Value: values[i],
			Key:   equation[1],
		})

		mp[equation[1]] = append(mp[equation[1]], Item{
			Value: 1 / values[i],
			Key:   equation[0],
		})
	}

	r := make([]float64, len(queries))
	for i, q := range queries {
		v, ok := dfs(q[0], q[1], len(set), mp, set)
		if !ok {
			r[i] = -1
			continue
		}

		r[i] = v
	}

	return r
}

func dfs(a string, b string, ttl int, mp map[string][]Item, set map[string]bool) (float64, bool) {
	if ttl < 0 {
		return 0, false
	}

	if !set[a] || !set[b] {
		return 0, false
	}

	if a == b {
		return 1, true
	}

	itemList, ok := mp[a]
	if !ok {
		return 0, false
	}

	for _, item := range itemList {
		r, ok := dfs(item.Key, b, ttl-1, mp, set)
		if ok {
			return r * item.Value, true
		}
	}

	return 0, false
}
