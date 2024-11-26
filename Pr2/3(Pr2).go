package main

import "fmt"

func Num3() {
	var year int
	fmt.Print("Введите год.")
	fmt.Scan(&year)
	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}
