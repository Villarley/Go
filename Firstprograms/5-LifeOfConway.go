package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	width, height := 25, 5
	board := createBoard(width, height)
	setRandomInitialState(board)

	for { // Bucle infinito
		printBoard(board)
		fmt.Println()
		updateBoard(board)
		time.Sleep(500 * time.Millisecond) // Pausa de medio segundo entre generaciones
	}
}

func setRandomInitialState(board [][]bool) {
	rand.Seed(time.Now().UnixNano())
	for y := range board {
		for x := range board[y] {
			board[y][x] = rand.Float64() < 0.5 // 50% de posibilidades de estar vivo
		}
	}
}
func createBoard(width, height int) [][]bool {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	return board
}
func printBoard(board [][]bool) {
	for _, row := range board {
		for _, cell := range row {
			if cell {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
func copyBoard(original [][]bool) [][]bool {
	height := len(original)
	width := len(original[0])
	copy := make([][]bool, height)
	for i := range original {
		copy[i] = make([]bool, width)
		for j := range original[i] {
			copy[i][j] = original[i][j]
		}
	}
	return copy
}
func updateBoard(board [][]bool) {
	height := len(board)
	width := len(board[0])
	tempBoard := copyBoard(board)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			liveNeighbors := countLiveNeighbors(tempBoard, x, y)

			if tempBoard[y][x] { // Si la celda está viva
				// Supervivencia o muerte
				board[y][x] = liveNeighbors == 2 || liveNeighbors == 3
			} else { // Si la celda está muerta
				// Nacimiento
				board[y][x] = liveNeighbors == 3
			}
		}
	}
}

func countLiveNeighbors(board [][]bool, x, y int) int {
	liveNeighbors := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				// Salta la propia celda
				continue
			}

			nx, ny := x+dx, y+dy
			// Comprueba si el vecino está dentro de los límites del tablero y está vivo
			if nx >= 0 && nx < len(board[0]) && ny >= 0 && ny < len(board) && board[ny][nx] {
				liveNeighbors++
			}
		}
	}
	return liveNeighbors
}
