package Pr2

import "fmt"

func Num11() {
	var (
		a int
		b int
	)
	fmt.Println("Введите диапазон(два числа через пробел) в котором следует найти простые числа.")
	fmt.Scanln(&a, &b)
	if a%2 == 0 {
		a++
	}
	for i := a; i <= b; i += 2 {
		counter := 0
		for j := 2; j <= i/2+1; j++ {
			if i%j == 0 {
				counter++
			}
		}
		if counter == 0 {
			fmt.Print(i, " ")
		}
	}
}
