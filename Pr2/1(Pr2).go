package main

import "fmt"

func Num1() {
	var (
		v0, a, t float64
	)
	fmt.Println("Введите скорость, ускорение, время")
	fmt.Scanln(&v0, &a, &t)
	if v0 < 0 {
		panic("Скорость не может быть отрицательной!")
	}
	if a == 0 {
		panic("Ускорение не может быть равно нулю!")
	}
	s := v0*t + a*t*t/2
	fmt.Printf("Текущее положение тела: %d", s)
}
