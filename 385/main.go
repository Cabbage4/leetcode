package main

import "strconv"

func main() {
}

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool {}

func (n NestedInteger) GetInteger() int {}

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger {}

func deserialize(s string) *NestedInteger {
	res := &NestedInteger{}
	if len(s) == 0 || (s[0] == '[' && s[1] == ']') {
		return res
	}
	if s[0] != '[' {
		val, _ := strconv.Atoi(s)
		res.SetInteger(val)
	} else {
		s := s[1 : len(s)-1]
		last := 0
		count := 0
		for i := 0; i <= len(s); i++ {
			if (i == len(s) || s[i] == ',') && count == 0 {
				res.Add(*(deserialize(s[last:i])))
				last = i + 1
			} else if s[i] == '[' {
				count++
			} else if s[i] == ']' {
				count--
			}
		}
	}
	return res
}
