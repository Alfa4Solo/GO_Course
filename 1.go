package main

import (
	"fmt"
	"strconv"
)

func Num1() {
	var (
		a string
		n int
		k int
	)
	fmt.Print("Введите число, исходную и конечную системы счисления.")
	fmt.Scan(&a)
	fmt.Scan(&n)
	fmt.Scan(&k)
	b, _ := strconv.ParseInt(a, n, 32)
	c := strconv.FormatInt(b, k)
	fmt.Println(c)
}
