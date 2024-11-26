package main

import (
	"fmt"
	"math"
	"os"
)

func Num4() {
	var (
		a, b, c float64
		D       float64
	)
	fmt.Print("Введите коэффиценты квадратного уравнения: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)
	fmt.Fscan(os.Stdin, &c)
	D = b*b - 4*a*c
	if D == 0 {
		fmt.Println(-b / (2 * a))
	} else if D > 0 {
		fmt.Println((-b + math.Sqrt(D)) / (2 * a))
		fmt.Println((-b - math.Sqrt(D)) / (2 * a))
	} else {
		fmt.Printf("%v + %v*i\n", -b/(2*a), math.Sqrt(-D)/(2*a))
		fmt.Printf("%v - %v*i\n", -b/(2*a), math.Sqrt(-D)/(2*a))
	}
}
