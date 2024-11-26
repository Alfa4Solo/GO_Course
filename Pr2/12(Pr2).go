package main

import (
	"fmt"
	"math/rand"
)

func Num12() {
	var (
		a, randRes int
	)
	fmt.Print("Попытайтесь угадать число: ")
	i := 0
	randRes = rand.Intn(100)
	for i == 0 {
		fmt.Scan(&a)
		if a > randRes {
			fmt.Println("Загаданное число меньше)")
		} else if a < randRes {
			fmt.Println("Загаданное число больше)")
		} else {
			fmt.Printf("Вы угадали! Это число %d", randRes)
			break
		}

	}

}
