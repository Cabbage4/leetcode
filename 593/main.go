package main

import "fmt"

func main() {
	data := [][]int{{0, 1}, {1, 1}, {1, 1}, {1, 0}}
	fmt.Println(validSquare(data[0], data[1], data[2], data[3]))
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] ||
		p1[0] == p3[0] && p1[1] == p3[1] ||
		p1[0] == p4[0] && p1[1] == p4[1] ||
		p2[0] == p3[0] && p2[1] == p3[1] ||
		p2[0] == p4[0] && p2[1] == p4[1] ||
		p3[0] == p4[0] && p3[1] == p4[1] {
		return false
	}

	return h(p1, p2, p3, p4) ||
		h(p1, p3, p2, p4) ||
		h(p1, p4, p2, p3) ||
		h(p2, p3, p1, p4) ||
		h(p2, p4, p1, p2) ||
		h(p3, p4, p1, p2)
}

func h(s []int, e []int, p1 []int, p2 []int) bool {

	s1x, s1y := p1[0]-s[0], p1[1]-s[1]
	s2x, s2y := p2[0]-s[0], p2[1]-s[1]

	if s1x*s2x+s1y*s2y != 0 {
		return false
	}

	e1x, e1y := p1[0]-e[0], p1[1]-e[1]
	e2x, e2y := p2[0]-e[0], p2[1]-e[1]

	if e1x*e2x+e1y*e2y != 0 {
		return false
	}

	a := s1x*s1x + s1y*s1y
	b := s2x*s2x + s2y*s2y
	c := e1x*e1x + e1y*e1y
	d := e2x*e2x + e2y*e2y

	if a != 0 && a == b && b == c && c == d {
		return true
	}

	return false
}
