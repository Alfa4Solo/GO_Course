package main

import (
	"awesomeProject/Pr1"
	"awesomeProject/Pr2"
	"awesomeProject/Pr3"
	"fmt"
)

func main() {
	var num int
	fmt.Println("С какой практической вы хотите поработать?")
	fmt.Scan(&num)
	switch num {
	case 1:
		Pr1.Practic1()
	case 2:
		Pr2.Practic2()
	case 3:
		Pr3.Practic3()
	}
}
