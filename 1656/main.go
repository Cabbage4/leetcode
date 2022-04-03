package main

func main() {
}

type OrderedStream struct {
	data []string
	ptr  int
}

func Constructor(n int) OrderedStream {
	r := OrderedStream{
		data: make([]string, n+1),
		ptr:  1,
	}
	return r
}

func (o *OrderedStream) Insert(k int, v string) []string {
	o.data[k] = v
	if k != o.ptr {
		return []string{}
	}

	r := make([]string, 0)
	for o.ptr < len(o.data) && o.data[o.ptr] != "" {
		r = append(r, o.data[o.ptr])
		o.ptr++
	}

	return r
}
