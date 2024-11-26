package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Num2() {
	fmt.Print("Введите выражение в польской записи.")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	s := in.Text()
	operands := "+-/*"
	arr := []int{}
	for _, val := range strings.Split(s, " ") {
		if !strings.Contains(operands, val) {
			conv, err := strconv.Atoi(val)
			if err != nil {
				panic("Неверная запись!")
			}
			arr = append(arr, conv)
		} else {
			switch val {
			case "+":
				{
					arr[len(arr)-2] = arr[len(arr)-2] + arr[len(arr)-1]
				}
			case "-":
				{
					arr[len(arr)-2] = arr[len(arr)-2] - arr[len(arr)-1]
				}
			case "*":
				{
					arr[len(arr)-2] = arr[len(arr)-2] * arr[len(arr)-1]
				}
			case "/":
				{
					arr[len(arr)-2] = arr[len(arr)-2] / arr[len(arr)-1]
				}
			}
			arr = append(arr[:len(arr)-2], arr[len(arr)-2])
		}
	}
	fmt.Println(arr[0])
}
