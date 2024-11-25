package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Num3() {
	var (
		a      []string
		sorted []int
	)
	fmt.Println("Введите в одну строчку массив эллементов.")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a = strings.Fields(s)
	for i := 0; i < len(a); i++ {
		p, _ := strconv.Atoi(a[i])
		sorted = append(sorted, p)
	}
	fmt.Println("Array of elements:", a)

	for i := 0; i < len(sorted); i++ {
		mi := sorted[i]
		Ad := i
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j] < mi {
				mi = sorted[j]
				Ad = j
			}
		}
		sorted[Ad] = sorted[i]
		sorted[i] = mi

	}
	fmt.Println("Sorted array of elements:", sorted)
}
