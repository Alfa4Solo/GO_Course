package Pr2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Practic2() {
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
		var n int
		fmt.Println("Введите номер символа до которого вывести последовательность Фибоначчи.")
		fmt.Scan(&n)
		var a, b []int
		a = append(a, 1, 1)
		b = Fib(a, n)
		for i := 2; i < n; i++ {
			a = append(a, a[i-1]+a[i-2])
		}
		fmt.Println(a)
		fmt.Println(b)
	case 7:
		Num7()
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
		var n int
		fmt.Scan(&n)
		s := ""
		for i := 0; i < 1; {
			s += string(rune(n % 2))
			n /= 2
			if n == 0 {
				i++
			}
		}
		fmt.Println(Reverse(s))
	case 15:
		fmt.Println("Введите значения клеток на поле 10x10 в формате '1','0'.")
		var (
			field [][]bool
			n     int
		)
		scanner := bufio.NewScanner(os.Stdin)
		for i := 0; i < height; i++ {
			scanner.Scan()
			line := scanner.Text()
			var row []bool
			for _, char := range strings.Split(line, "") {
				if char == "1" {
					row = append(row, true)
				} else if char == "0" {
					row = append(row, false)
				}
			}
			field = append(field, row)
		}
		fmt.Println("Введите количество шагов для симуляции.")
		fmt.Scan(&n)
		Num15(field, n)
	}
}
