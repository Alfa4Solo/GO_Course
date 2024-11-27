package Pr1

import (
	"fmt"
	"math"
)

func Num6() {

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
	} else if (ch == "/") && (b != 0) {
		fmt.Printf("Результат операции \"%s\": %.3f", ch, a/b)
	} else if ch == "^" && (a < 0) {
		fmt.Printf("Результат операции \"%s\": %.3f", ch, math.Pow(a, b))
	} else if ch == "%" {
		fmt.Printf("Результат операции \"%s\": %.f", ch, math.Mod(a, b))
	}
}
