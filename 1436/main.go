package main

func main() {
}

func destCity(paths [][]string) string {
	pmp := make(map[string]bool)
	for _, p := range paths {
		pmp[p[1]] = true
	}

	for _, p := range paths {
		delete(pmp, p[0])
	}

	for k := range pmp {
		return k
	}

	return ""
}
