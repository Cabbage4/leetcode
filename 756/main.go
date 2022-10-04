package main

import "fmt"

func main() {
	fmt.Println(pyramidTransition("DAAAAD", []string{"DAD", "DAE", "DAB", "DAF", "DAC", "EAD", "EAE", "EAB", "EAF", "EAC", "BAD", "BAE", "BAB", "BAF", "BAC", "FAD", "FAE", "FAB", "FAF", "FAC", "CAD", "CAE", "CAB", "CAF", "CAC", "ADD", "ADE", "ADB", "ADF", "ADC", "AED", "AEE", "AEB", "AEF", "AEC", "ABD", "ABE", "ABB", "ABF", "ABC", "AFD", "AFE", "AFB", "AFF", "AFC", "ACD", "ACE", "ACB", "ACF", "ACC", "AAD", "AAE", "AAB", "AAF", "AAC", "AAA"}))
	fmt.Println(pyramidTransition("AABA", []string{"AAB", "AAC", "BCD", "BBE", "DEF"}))
	fmt.Println(pyramidTransition("BCD", []string{"BCC", "CDE", "CEA", "FFF"}))
}

func pyramidTransition(bottom string, allowed []string) bool {
	mp := make(map[string][]byte)
	for _, a := range allowed {
		mp[a[:2]] = append(mp[a[:2]], a[2])
	}

	var dfs func(string, string) bool
	dfs = func(b string, t string) bool {
		if len(b) == 1 {
			return true
		}

		if index := len(t); index < len(b)-1 {
			for _, v := range mp[b[index:index+2]] {
				tmp := t + string(v)
				if len(tmp)+1 == len(b) && dfs(tmp, "") {
					return true
				}

				if dfs(b, tmp) {
					return true
				}
			}
		}

		return false
	}

	return dfs(bottom, "")
}
