package Pr3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Practic3() {
	var n int
	fmt.Println("Введите номер задачи, которую необходимо решить.")
	fmt.Scan(&n)
	switch n {
	case 2:
		Num2()
	case 3:
		Num3()
	case 4:

		fmt.Println("Введите массив слов.")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		s = s[2 : len(s)-3]
		elements := strings.Split(s, "\", \"")
		var mas []string
		for _, elem := range elements {
			mas = append(mas, elem)
		}
		fmt.Println(mas)
		fmt.Println("Самый длинный общий префикс: ", Num4(mas))
	}
}
