package Pr3

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

func Num2() {
	fmt.Print("Введите выражение. ")
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	line = line[:len(line)-1]
	s := strings.Fields(line)
	elems := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "(") || strings.Contains(s[i], ")") {
		}
		elems[i] = s[i]
	}
	//operands := "+-/*"
	//arr := []int{}
	/*for _, val := range strings.Split(line, " ") {
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
	fmt.Println(strings.Split(line, " "))
}
