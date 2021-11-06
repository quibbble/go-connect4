package go_connect4

import (
	"encoding/json"
	bg "github.com/quibbble/go-boardgame"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Builder_BGN(t *testing.T) {
	builder := Builder{}
	teams := []string{TeamA, TeamB}
	connect4, err := builder.CreateWithBGN(&bg.BoardGameOptions{Teams: teams, Seed: 123})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = connect4.Do(&bg.BoardGameAction{
		Team:       TeamB,
		ActionType: ActionPlaceDisk,
		MoreDetails: PlaceDiskActionDetails{
			Column: 1,
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = connect4.Do(&bg.BoardGameAction{
		Team:       TeamA,
		ActionType: ActionPlaceDisk,
		MoreDetails: PlaceDiskActionDetails{
			Column: 1,
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	game := connect4.GetBGN()
	connect4Loaded, err := builder.Load(game)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected, _ := json.Marshal(connect4)
	actual, _ := json.Marshal(connect4Loaded)
	assert.Equal(t, expected, actual)
}
