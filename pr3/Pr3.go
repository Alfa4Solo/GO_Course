package Pr3

import (
	"fmt"
)

func Practic3() {
	var n int
	fmt.Println("Введите номер задачи, которую необходимо решить.")
	fmt.Scan(&n)
	switch n {
	case 2:
		Num2()
	case 3:
		Num3()
	case 4:
		//Num4()
	}
}
