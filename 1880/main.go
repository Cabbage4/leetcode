package main

import "fmt"

func main() {
	fmt.Println(isSumEqual("jjbcgj", "hddda", "bagefjj"))
	fmt.Println(isSumEqual("j", "j", "bi"))
	fmt.Println(isSumEqual("acb", "cba", "cdb"))
}

func isSumEqual(s1 string, s2 string, s3 string) bool {
	v1 := getV(s1)
	v2 := getV(s2)
	v3 := getV(s3)

	if len(v1) < len(v2) {
		v1, v2 = v2, v1
	}

	var sw byte = 0
	v1i := len(v1) - 1
	for v2i := len(v2) - 1; v2i >= 0; v2i-- {
		t := v1[v1i] - '0' + v2[v2i] - '0' + sw
		sw = t / 10
		v1[v1i] = t%10 + '0'
		v1i--
	}

	if sw != 0 {
		if v1i == -1 {
			v1 = append([]byte{sw + '0'}, v1...)
		} else {
			for ; v1i >= 0; v1i-- {
				t := v1[v1i] - '0' + sw
				sw = t / 10
				v1[v1i] = t%10 + '0'
				if sw == 0 {
					break
				}
			}

			if sw != 0 {
				v1 = append([]byte{sw + '0'}, v1...)
			}
		}
	}

	if len(v1) != len(v3) {
		return false
	}

	for i := 0; i < len(v1); i++ {
		if v1[i] != v3[i] {
			return false
		}
	}

	return true
}

func getV(s string) []byte {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = '0' + s[i] - 'a'
	}

	i := 0
	for ; i < len(r); i++ {
		if r[i] != '0' {
			break
		}
	}

	return r[i:]
}
