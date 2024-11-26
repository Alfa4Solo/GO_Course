package main

import (
	"fmt"
)

func Num9() {
	var (
		a, b int
	)
	fmt.Println("Введите диапазон(два числа через пробел) в котором будет выведена таблица умножения")
	fmt.Scanln(&a, &b)
	for j := 1; j <= 10; j++ {
		for i := a; i <= b; i++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}
