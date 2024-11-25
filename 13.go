package main

import (
	"fmt"
	"math"
)

func Num13() {
	var (
		a int
		b int
	)
	fmt.Println("Введите диапазон(два числа через пробел) в котором следует найти числа Армстронга.")
	fmt.Scanln(&a, &b)
	var results []int
	for i := a; i <= b; i++ {
		size := int(math.Log10(float64(i))) + 1
		sum := 0

		for j := i; j > 0; j /= 10 {
			x := j % 10
			sum += int(math.Pow(float64(x), float64(size)))
		}
		if sum == i {
			results = append(results, i)
		}
	}
	fmt.Println(results)
}
