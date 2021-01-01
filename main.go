package main

import (
	"fmt"
	"strconv"
)

var board = make([]string, 9, 9)
var currPlayer string

func main() {
	initialSetup()
	play()
}

func initialSetup() {
	// Filling up the board
	for i := 0; i < len(board); i ++ {
		board[i] = strconv.Itoa(i+1) 
	}
	//Asking the user to choose a symbol
	for (currPlayer != "X") && (currPlayer != "O") {
		fmt.Println("Please choose X or O")
		fmt.Scan(&currPlayer)
	}
	displayBoard()
}

func play() {
	fmt.Printf("Player %v please choose a position to place your move\n", currPlayer)
	var position string
	for ;; {
		fmt.Scan(&position)
		if isValidPos(position) {
			break
		}else {
			fmt.Println("Please choose a valid position")
		}
	}
	// Placing the chosen move on the board
	boardIndex, _ := strconv.Atoi(position)
	boardIndex -= 1
	board[boardIndex] = currPlayer
	displayBoard()
	if (gameOver()) {
		return
	}
	// Switch player
	if (currPlayer == "X") {
		currPlayer = "O"
	}else {
		currPlayer = "X"
	}
	play()
}

func displayBoard() {
	fmt.Println("      ||     ||      ")
	fmt.Printf("   %v  ||  %v  ||  %v   \n",board[0], board[1], board[2])
	fmt.Println("=====================")
	fmt.Printf("   %v  ||  %v  ||  %v   \n",board[3], board[4], board[5])
	fmt.Println("=====================")
	fmt.Printf("   %v  ||  %v  ||  %v   \n",board[6], board[7], board[8])
	fmt.Println("      ||     ||      ")
}

func isValidPos(position string) bool{
	for _, v := range board {
		if position == v {
			return true
		}
	}
	return false
}

func gameOver() bool{
	// Rowwise Check
	for r := 0; r < 9; r += 3 {
		if (board[r] == board[r+1]) && (board[r] == board[r+2]) {
			fmt.Printf("Player %v has won the game\n", currPlayer)
			return true
		}
	}
	// Columnwise Check
	for c := 0; c < 3; c += 1 {
		if (board[c] == board[c+3]) && (board[c] == board[c+6]) {
			fmt.Printf("Player %v has won the game\n", currPlayer)
			return true
		}
	}
	// Diagonal 1 Check
	if (board[0] == board[4]) && (board[0] == board[8]) {
		fmt.Printf("Player %v has won the game\n", currPlayer)
		return true
	}
	// Diagonal 2 Check
	if (board[2] == board[4]) && (board[2] == board[6]) {
		fmt.Printf("Player %v has won the game\n", currPlayer)
		return true
	}
	// All positions filled
	for _, v := range board {
		if v != "X" && v != "O" {
			return false
		}
	}
	fmt.Println("Game draw")
	return true
}