package Pr3

import (
	"fmt"
	"slices"
)

func Num3() {
	var (
		n int
	)
	fmt.Println("Введите количество строк треугольника Паскаля")
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println("1")
	} else if n > 1 {
		fmt.Println("1")
		fmt.Println("1 1")
		Pascal := []int{1, 1}
		for i := 0; i < n-2; i++ {
			line := []int{}
			for j := 0; j < len(Pascal)-1; j++ {
				line = append(line, Pascal[j]+Pascal[j+1])
			}
			Pascal = []int{1, 1}
			Pascal = slices.Insert(Pascal, 1, line...)
			for k := 0; k < len(Pascal); k++ {
				fmt.Printf("%d ", Pascal[k])
			}
			fmt.Println()
		}
	} else {
		fmt.Print(0)
	}

}
