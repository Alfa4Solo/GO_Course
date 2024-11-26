package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Num9() {
	var (
		ma int
		w  string
	)
	fmt.Println("Введите строку с словами")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	words := strings.Fields(s)
	for _, word := range words {
		ma = len(words[0])
		if ma < len(word) {
			ma = len(word)
			w = word
		}
	}
	fmt.Print(w)
}
