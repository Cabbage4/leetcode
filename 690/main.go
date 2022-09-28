package main

func main() {
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mp := make(map[int]*Employee)
	for _, e := range employees {
		mp[e.Id] = e
	}

	var r int
	var dfs func(int)
	dfs = func(id int) {
		r += mp[id].Importance

		for _, s := range mp[id].Subordinates {
			dfs(mp[s].Id)
		}
	}

	dfs(id)

	return r
}
