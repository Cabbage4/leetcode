package main

func main() {

}

func findTheDifference(s string, t string) byte {
	var c int32 = 0
	for _, v := range t {
		c = c ^ v
	}

	for _, v := range s {
		c = c ^ v
	}

	return byte(c)
}
