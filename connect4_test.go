package go_connect4

import (
	"testing"

	bg "github.com/quibbble/go-boardgame"
	"github.com/stretchr/testify/assert"
)

const (
	TeamA = "TeamA"
	TeamB = "TeamB"
)

func Test_Connect4(t *testing.T) {
	connect4, err := NewConnect4(&bg.BoardGameOptions{
		Teams: []string{TeamA, TeamB},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	connect4.state.turn = TeamA

	// place disk in column 0
	err = connect4.Do(&bg.BoardGameAction{
		Team:       TeamA,
		ActionType: ActionPlaceDisk,
		MoreDetails: PlaceDiskActionDetails{
			Column: 0,
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(t, TeamA, *connect4.state.board.board[rows-1][0])
}
