package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(largestInteger(1234))
}

type Heap struct {
	o  []int
	oi int
	e  []int
	ei int
}

func (h *Heap) Push(x int) {
	if x%2 == 0 {
		h.e[x]++
		if x > h.ei {
			h.ei = x
		}
		return
	}
	h.o[x]++
	if x > h.oi {
		h.oi = x
	}
}

func (h *Heap) PopO() int {
	r := h.oi
	h.o[h.oi]--
	for h.oi >= 0 && h.o[h.oi] == 0 {
		h.oi--
	}
	return r
}

func (h *Heap) PopE() int {
	r := h.ei
	h.e[h.ei]--
	for h.ei >= 0 && h.e[h.ei] == 0 {
		h.ei--
	}
	return r
}

func largestInteger(num int) int {
	h := &Heap{
		o: make([]int, 10),
		e: make([]int, 10),
	}

	b := []byte(strconv.FormatInt(int64(num), 10))
	for num > 0 {
		h.Push(num % 10)
		num = num / 10
	}

	for i := 0; i < len(b); i++ {
		if (b[i]-'0')%2 == 0 {
			b[i] = byte(h.PopE()) + '0'
			continue
		}
		b[i] = byte(h.PopO()) + '0'
	}

	r, _ := strconv.ParseInt(string(b), 10, 32)
	return int(r)
}
