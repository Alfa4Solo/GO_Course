package Pr1

import "fmt"

func Num11() {
	var n int
	fmt.Println("Введите номер символа до которого вывести последовательность Фибоначчи.")
	fmt.Scan(&n)
	var a []int
	a = append(a, 1, 1)
	for i := 2; i < n; i++ {
		a = append(a, a[i-1]+a[i-2])
	}
	fmt.Println(a)
}
