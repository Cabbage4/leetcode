package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(1994))
	//fmt.Println(intToRoman(3))
	//fmt.Println(intToRoman(58))
}

func intToRoman(num int) string {
	l := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	l0 := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	l00 := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	l000 := []string{"", "M", "MM", "MMM"}
	return l000[num/1000] + l00[(num%1000)/100] + l0[(num%100)/10] + l[num%10]
}
