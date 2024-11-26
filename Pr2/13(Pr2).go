package main

import "fmt"

func Num13() {
	var n int
	fmt.Print("Введите высоту ромба.")
	fmt.Scan(&n)
	s := "*"
	for i := 1; i < n; i += 2 {
		spaces := ""
		for j := 0; j < n-i; j++ {
			spaces += " "
		}
		fmt.Println(spaces, s)
		s += "**"
	}
	for i := n; i >= 1; i -= 2 {
		spaces := ""
		for j := 0; j < n-i; j++ {
			spaces += " "
		}
		fmt.Println(spaces, s)
		if len(s) > 2 {
			s = s[:len(s)-2]
		}
	}
}
