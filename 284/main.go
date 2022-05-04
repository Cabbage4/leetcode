package main

func main() {
}

type Iterator interface {
	hasNext() bool
	next() int
}

type PeekingIterator struct {
	iter Iterator
	list []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: *iter,
		list: make([]int, 0),
	}
}

func (p *PeekingIterator) hasNext() bool {
	if len(p.list) != 0 {
		return true
	}
	return p.iter.hasNext()
}

func (p *PeekingIterator) next() int {
	if len(p.list) != 0 {
		v := p.list[0]
		p.list = p.list[1:]
		return v
	}
	return p.iter.next()
}

func (p *PeekingIterator) peek() int {
	if len(p.list) != 0 {
		return p.list[0]
	}

	v := p.iter.next()
	p.list = append(p.list, v)
	return v
}
