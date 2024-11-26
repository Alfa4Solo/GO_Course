package main

import (
	"fmt"
)

func Num5() {

	var (
		a  float64
		b  float64
		ch string
	)
	fmt.Println("Введите два числа, а затем операцию между ними.")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scanf("%s", &ch)
	if ch == "+" {
		fmt.Printf("Результат операции \"%s\": %.3f", ch, a+b)
	} else if ch == "-" {
		fmt.Printf("Результат операции \"%s\": %.3f", ch, a-b)
	} else if ch == "*" {
		fmt.Printf("Результат операции \"%s\": %.3f", ch, a*b)
	} else if ch == "/" {
		if b == 0 {
			panic("Делить на ноль нельзя!")
		}
		fmt.Printf("Результат операции \"%s\": %.3f", ch, a/b)
	}
}
