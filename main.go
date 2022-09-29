package main

import "fmt"

func main() {
	m := Constructor()
	fmt.Println(m.Book(10, 20))
	fmt.Println(m.Book(50, 60))
	fmt.Println(m.Book(10, 40))
	fmt.Println(m.Book(5, 15))
	fmt.Println(m.Book(5, 10))
	fmt.Println(m.Book(25, 55))
}

type MyCalendarTwo struct {
}

func Constructor() MyCalendarTwo {
}

func (m *MyCalendarTwo) Book(s int, e int) bool {
}
