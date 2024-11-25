package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Println("Введите номер задачи, которую необходимо решить.")
	fmt.Scan(&n)
	switch n {
	case 1:
		Num1()
	case 2:
		Num2()
	case 3:
		Num3()
	case 4:
		Num4()
	case 5:
		Num5()
	case 6:
		Num6()
	case 7:
		fmt.Println("Введите строку для проверки на паллиндром.")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		result := IsPalindrome(s)
		fmt.Printf("Строка: \"%s\" является палиндромом: %v\n", s, result)
	case 8:
		Num8()
	case 9:
		Num9()
	case 10:
		Num10()
	case 11:
		Num11()
	case 12:
		Num12()
	case 13:
		Num13()
	case 14:
		Num14()
	case 15:
		Num15()
	}
}
