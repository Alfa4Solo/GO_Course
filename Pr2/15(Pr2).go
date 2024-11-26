package main

import (
	"fmt"
	"time"
)

const (
	width  = 10 // Ширина поля
	height = 10 // Высота поля
)

func countNeighbors(board [][]bool, x, y int) int {
	neighbors := 0
	positions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, pos := range positions {
		nx, ny := x+pos[0], y+pos[1]
		if nx >= 0 && nx < width && ny >= 0 && ny < height && board[nx][ny] {
			neighbors++
		}
	}
	return neighbors
}

func Num15(board [][]bool, steps int) {

	for i := 0; i < steps; i++ {
		fmt.Print("\033[H\033[2J")
		for _, row := range board {
			fmt.Print("|")
			for _, cell := range row {
				if cell {
					fmt.Print("█") // Живая клетка
				} else {
					fmt.Print(" ") // Мертвая клетка
				}
			}
			fmt.Println("|")
			time.Sleep(300 * time.Millisecond)
			newBoard := make([][]bool, width)
			for i := range newBoard {
				newBoard[i] = make([]bool, height)
			}

			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					neighbors := countNeighbors(board, x, y)

					if board[x][y] {
						if neighbors == 2 || neighbors == 3 {
							newBoard[x][y] = true
						} else {
							newBoard[x][y] = false
						}
					} else {
						if neighbors == 3 {
							newBoard[x][y] = true
						}
					}
				}
			}
			board = newBoard
		}
		fmt.Println("____________")
	}
}
