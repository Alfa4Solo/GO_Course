package main

import (
	"fmt"
	//"strings"
)

func Num5() {

	var (
		str  string
		str1 string
	)
	fmt.Println("Введите строку и потенциальную подстроку.")
	fmt.Scanf("%s\n", &str)
	fmt.Scanf("%s\n", &str1)
	fst := -1
	for i := 0; i < len(str)-len(str1)+1; i++ {
		if str[i] == str1[0] && fst == -1 {
			for j := 1; j < len(str1)+1; j++ {
				if j == len(str1) {
					fst = i
					break
				} else if str[i+j] != str1[j] {
					break
				}
			}

		}
	}
	fmt.Print(fst)
	//fmt.Printf("Строка \"%s\" содержться в \"%s\": %t", str, str1, strings.Contains(str, str1))
}
