package main

import "fmt"

func main() {
	fmt.Println(totalMoney(4))
	//fmt.Println(totalMoney(10))
}

func totalMoney(n int) int {
	f := n / 7
	d := n % 7
	return (49+7*f)*f/2 + (2*f+d+1)*d/2
}
