package Pr2

import "fmt"

func Num7() {
	var (
		a int
	)
	fmt.Println("Введите число")
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println("не простое")
		return
	}
	for i := 3; i <= a/2+1; i += 2 {
		if a%i == 0 {
			fmt.Println("не простое")
			return
		}
	}
	fmt.Println("простое")
}
