package Pr3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Num2() {
	fmt.Print("Введите выражение. ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	operands := "+-/*()"
	//arr := []int{}
	/*for _, val := range strings.Split(s, " ") {
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
	}*/
	fmt.Println(strings.Split(s, operands))
}
