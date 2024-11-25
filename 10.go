package main

import "fmt"

func Num10() {
	var year int
	fmt.Println("Введите год.")
	fmt.Scan(&year)
	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}
