package Pr1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Num4() {
	var (
		mas1 []int
		mas2 []int
		mas  []int
	)
	fmt.Println("Введите в одну строчку массив эллементов#1.")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a := strings.Fields(s)
	for i := 0; i < len(a); i++ {
		p, _ := strconv.Atoi(a[i])
		mas1 = append(mas1, p)
	}
	fmt.Println("Введите в одну строчку массив эллементов#2.")
	s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	a = strings.Fields(s)

	for i := 0; i < len(a); i++ {
		p, _ := strconv.Atoi(a[i])
		mas2 = append(mas2, p)
	}
	fmt.Println("First array of elements:", mas1)
	fmt.Println("Second array of elements:", mas2)
	mas = append(mas1, mas2...)
	for i := 0; i < len(mas); i++ {
		mi := mas[i]
		Ad := i
		for j := i + 1; j < len(mas); j++ {
			if mas[j] < mi {
				mi = mas[j]
				Ad = j
			}
		}
		mas[Ad] = mas[i]
		mas[i] = mi

	}
	fmt.Println("Sorted array of elements:", mas)
}
