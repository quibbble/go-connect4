package go_connect4

import "fmt"

const (
	rows    = 6
	columns = 7
)

type board struct {
	board [][]*string // 0,0 is top left corner
}

func newBoard() *board {
	var b = make([][]*string, rows)
	for row := 0; row < rows; row++ {
		b[row] = make([]*string, columns)
	}
	return &board{
		board: b,
	}
}

func (b *board) PlaceDisk(player string, col int) error {
	for row := len(b.board) - 1; row >= 0; row-- {
		if b.board[row][col] == nil {
			b.board[row][col] = &player
			return nil
		}
	}
	return fmt.Errorf("column %d is full", col)
}

func (b *board) isFull() bool {
	for _, row := range b.board {
		for _, it := range row {
			if it == nil {
				return false
			}
		}
	}
	return true
}
