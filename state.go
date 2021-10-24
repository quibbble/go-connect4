package go_connect4

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"math/rand"
)

type state struct {
	turn    string
	teams   []string
	winners []string
	board   *board
}

func newState(teams []string) *state {
	return &state{
		turn:    teams[rand.Intn(len(teams))],
		teams:   teams,
		winners: make([]string, 0),
		board:   newBoard(),
	}
}

func (s *state) PlaceDisk(team string, column int) error {
	if len(s.winners) > 0 {
		return &bgerr.Error{
			Err:    fmt.Errorf("%s game already completed", key),
			Status: bgerr.StatusGameOver,
		}
	}
	if team != s.turn {
		return &bgerr.Error{
			Err:    fmt.Errorf("currently %s's turn", s.turn),
			Status: bgerr.StatusWrongTurn,
		}
	}
	if column < 0 || column > columns-1 {
		return &bgerr.Error{
			Err:    fmt.Errorf("column %d is out of bounds", column),
			Status: bgerr.StatusInvalidActionDetails,
		}
	}
	if err := s.board.PlaceDisk(team, column); err != nil {
		return &bgerr.Error{
			Err:    err,
			Status: bgerr.StatusInvalidAction,
		}
	}
	if s.board.isFull() {
		s.winners = s.teams
		return nil
	}
	if winner := findWinner(s.board); winner != nil {
		s.winners = []string{*winner}
		return nil
	}
	for idx, team := range s.teams {
		if team == s.turn {
			s.turn = s.teams[(idx+1)%len(s.teams)]
			break
		}
	}
	return nil
}

// nil if no winner, winner name if winner
func findWinner(board *board) *string {
	// check vertical
	for row := 0; row < rows-3; row++ {
		for col := 0; col < columns; col++ {
			if board.board[row][col] == nil || board.board[row+1][col] == nil || board.board[row+2][col] == nil || board.board[row+3][col] == nil {
				continue
			}
			player := *board.board[row][col]
			if player == *board.board[row+1][col] && player == *board.board[row+2][col] && player == *board.board[row+3][col] {
				return &player
			}
		}
	}
	// check horizontal
	for row := 0; row < rows; row++ {
		for col := 0; col < columns-3; col++ {
			if board.board[row][col] == nil || board.board[row][col+1] == nil || board.board[row][col+2] == nil || board.board[row][col+3] == nil {
				continue
			}
			player := *board.board[row][col]
			if player == *board.board[row][col+1] && player == *board.board[row][col+2] && player == *board.board[row][col+3] {
				return &player
			}
		}
	}
	// check positive diagonal
	for row := 0; row < rows-3; row++ {
		for col := 0; col < columns-3; col++ {
			if board.board[row][col] == nil || board.board[row+1][col+1] == nil || board.board[row+2][col+2] == nil || board.board[row+3][col+3] == nil {
				continue
			}
			player := *board.board[row][col]
			if player == *board.board[row+1][col+1] && player == *board.board[row+2][col+2] && player == *board.board[row+3][col+3] {
				return &player
			}
		}
	}
	// check negative diagonal
	for row := 3; row < rows; row++ {
		for col := 0; col < columns-3; col++ {
			if board.board[row][col] == nil || board.board[row-1][col+1] == nil || board.board[row-2][col+2] == nil || board.board[row-3][col+3] == nil {
				continue
			}
			player := *board.board[row][col]
			if player == *board.board[row-1][col+1] && player == *board.board[row-2][col+2] && player == *board.board[row-3][col+3] {
				return &player
			}
		}
	}
	return nil
}