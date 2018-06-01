package main

import (
	"errors"
	"fmt"
	"strings"
)

// Board ...
type Board struct {
	board [3][3]string
}

// NewBoard ...
func NewBoard() *Board {
	b := &Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.board[i][j] = ""
		}
	}
	return b
}

// String ...
func (b *Board) String() string {
	board := ""

	for i := 0; i < len(b.board); i++ {
		var row []string
		for j := 0; j < len(b.board[i]); j++ {
			val := b.board[i][j]
			if val == "" {
				val = "-"
			}
			row = append(row, val)
		}
		board += strings.Join(row, "|") + "\n"
	}

	return board
}

// Move ...
func (b *Board) Move(x, y int, player *Player) error {
	if b.board[x][y] != "" {
		return errors.New("Space is already occupied")
	}

	b.board[x][y] = fmt.Sprintf("%s", player)
	return nil
}

// AIMove ...
// This is just a generic "make a valid move" implementation.
func (b *Board) AIMove(player *Player) error {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board); j++ {
			if b.board[i][j] == "" {
				return b.Move(i, j, player)
			}
		}
	}

	return errors.New("No available moves")
}

// IsFull ...
func (b *Board) IsFull() bool {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			if b.board[i][j] == "" {
				return false
			}
		}
	}

	return true
}
