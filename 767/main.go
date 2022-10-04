package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reorganizeString("aabbcc"))
	fmt.Println(reorganizeString("vvvlo"))
	fmt.Println(reorganizeString("aaabbb"))
}

var mp [26]int

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return mp[h.IntSlice[i]] > mp[h.IntSlice[j]]
}

func (h *Heap) Pop() interface{} {
	r := h.IntSlice[h.IntSlice.Len()-1]
	h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1]
	return r
}

func (h *Heap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func reorganizeString(s string) string {
	mp = [26]int{}

	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
	}

	maxCount := 0
	h := Heap{[]int{}}
	for i := 0; i < len(mp); i++ {
		if mp[i] == 0 {
			continue
		}

		if maxCount < mp[i] {
			maxCount = mp[i]
		}

		heap.Push(&h, i)
	}

	if maxCount > (len(s)+1)/2 {
		return ""
	}

	r := bytes.NewBuffer(make([]byte, 0, len(s)))
	for len(h.IntSlice) > 1 {
		a, b := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if mp[a]--; mp[a] > 0 {
			heap.Push(&h, a)
		}
		if mp[b]--; mp[b] > 0 {
			heap.Push(&h, b)
		}
		r.WriteByte(byte(a + 'a'))
		r.WriteByte(byte(b + 'a'))
	}

	if h.IntSlice.Len() > 0 {
		r.WriteByte(byte(h.IntSlice[0] + 'a'))
	}

	return r.String()
}
