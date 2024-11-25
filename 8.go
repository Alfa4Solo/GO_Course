package main

import (
	"fmt"
)

func Num8() {
	var (
		a1, a2, b1, b2, a3, b3 int
	)
	fmt.Println("Введите координаты первого отрезка:")
	fmt.Scanln(&a1, &b1)
	fmt.Println("Введите координаты второго отрезка:")
	fmt.Scanln(&a2, &b2)
	fmt.Println("Введите координаты третьего отрезка:")
	fmt.Scanln(&a3, &b3)
	ma := max(a1, a2)
	mi := min(b1, b2)
	if (ma <= mi) && (max(ma, a3) <= min(mi, b3)) {
		fmt.Printf("Пересечение трех отрезков: %s\n", "true")
	} else {
		fmt.Printf("Пересечение трех отрезков: %s\n", "false")
	}
}
