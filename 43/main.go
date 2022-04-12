package main

import "fmt"

func main() {
	fmt.Println(multiply("408", "5"))
	fmt.Println(multiply("98", "9"))
	fmt.Println(multiply("123", "456"))
	fmt.Println(multiply("2", "3"))
}

func multiply(o1 string, o2 string) string {
	if o1 == "0" || o2 == "0" {
		return "0"
	}

	bit := 0
	p1 := make([]byte, 0)
	for i := len(o1) - 1; i >= 0; i-- {
		p2 := make([]byte, 0)
		var w1 byte
		for j := len(o2) - 1; j >= 0; j-- {
			t := (o1[i]-'0')*(o2[j]-'0') + w1
			w1 = t / 10
			p2 = append(p2, t%10)
		}
		if w1 != 0 {
			p2 = append(p2, w1)
		}

		p1i := bit
		p2i := 0

		tp := make([]byte, 0)
		for k := 0; k < p1i; k++ {
			tp = append(tp, p1[k])
		}

		var w2 byte
		for p2i < len(p2) && p1i < len(p1) {
			t := p1[p1i] + p2[p2i] + w2
			w2 = t / 10
			tp = append(tp, t%10)
			p1i++
			p2i++
		}

		if w2 != 0 {
			for p1i < len(p1) {
				t := p1[p1i] + w2
				w2 = t / 10
				tp = append(tp, t%10)
				p1i++
			}

			for p2i < len(p2) {
				t := p2[p2i] + w2
				w2 = t / 10
				tp = append(tp, t%10)
				p2i++
			}

			if w2 != 0 {
				tp = append(tp, w2)
			}
		} else {
			if p1i < len(p1) {
				tp = append(tp, p1[p1i:]...)
			}
			if p2i < len(p2) {
				tp = append(tp, p2[p2i:]...)
			}
		}

		p1 = tp
		bit++
	}

	r := make([]byte, 0)
	for i := len(p1) - 1; i >= 0; i-- {
		r = append(r, p1[i]+'0')
	}
	return string(r)
}
