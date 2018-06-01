package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	board := NewBoard()

	player1 := NewPlayer("X")
	player2 := NewPlayer("O")

	for {
		scanner.Scan()
		text := strings.TrimRight(scanner.Text(), " ")

		pos := strings.Split(text, " ")
		if len(pos) != 2 {
			fmt.Fprintf(os.Stderr, "Please enter a valid move of row col, like [0 1]\n")
			continue
		}

		row, err := strconv.Atoi(pos[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Please enter a valid row for the first number %v\n", err)
			continue
		}

		col, err := strconv.Atoi(pos[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Please enter a valid col for the first number %v\n", err)
			continue
		}

		if row < 0 || col < 0 || row > 2 || col > 2 {
			fmt.Fprintf(os.Stderr, "Please enter a number 0-2 inclusive\n")
			continue
		}

		if err := board.Move(row, col, player1); err != nil {
			fmt.Fprintf(os.Stderr, "Invalid move %v\n", err)
			continue
		}

		// Let the "computer" make a move.
		board.AIMove(player2)

		fmt.Fprintf(os.Stdout, "%s", board)

		if board.IsFull() {
			fmt.Println("Game over. GJ")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
