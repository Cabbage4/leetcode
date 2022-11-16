package main

import "math/rand"

func main() {
}

type Solution struct {
	x float64
	y float64
	r float64
}

func Constructor(radius float64, x float64, y float64) Solution {
	return Solution{
		x: x,
		y: y,
		r: radius,
	}
}

func (s *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y < 1 {
			return []float64{s.x + x*s.r, s.y + y*s.r}
		}
	}
}
