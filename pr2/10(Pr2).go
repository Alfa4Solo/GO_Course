package Pr2

import "fmt"

func Num10() {
	var (
		a int
		b int
	)
	fmt.Println("Введите два числа через пробел, чтобы определить их НОД.")
	fmt.Scanln(&a, &b)
	ma := 0
	if a%b == 0 || b%a == 0 {
		fmt.Println(min(a, b))
		return
	}
	for i := 2; i < min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			ma = i
		}
	}
	fmt.Println(ma)
}
