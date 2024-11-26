package main

import (
	"fmt"
	"math"
)

func Num8() {
	var (
		a int
	)
	fmt.Println("Введите число(проверка на Армстронга).")
	fmt.Scanln(&a)
	size := int(math.Log10(float64(a))) + 1
	sum := 0
	for j := a; j > 0; j /= 10 {
		x := j % 10
		sum += int(math.Pow(float64(x), float64(size)))
	}
	if sum == a {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
