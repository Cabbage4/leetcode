package main

import "fmt"

func main() {
	c := Constructor(1, 1, 0)
	fmt.Println(c.AddCar(1))
	fmt.Println(c.AddCar(2))
	fmt.Println(c.AddCar(3))
}

type ParkingSystem struct {
	list []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	r := new(ParkingSystem)
	r.list = make([]int, 4)
	r.list[1] = big
	r.list[2] = medium
	r.list[3] = small

	return *r
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.list[carType] > 0 {
		this.list[carType]--
		return true
	}

	return false
}
