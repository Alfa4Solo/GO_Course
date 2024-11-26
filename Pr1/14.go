package main

import (
	"fmt"
)

func Num14() {
	var (
		s         string
		sReversed string
	)
	fmt.Println("Введите строку")
	fmt.Scanf("%s", &s)
	for i := len(s) - 1; i > -1; i-- {
		sReversed = sReversed + string(s[i])
	}
	fmt.Println(sReversed)
}
