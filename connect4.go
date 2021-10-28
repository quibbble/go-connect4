package go_connect4

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
)

const (
	minTeams = 2
	maxTeams = 3
)

type Connect4 struct {
	state   *state
	actions []*bg.BoardGameAction
}

func NewConnect4(options bg.BoardGameOptions) (*Connect4, error) {
	if len(options.Teams) < minTeams {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("at least %d teams required to create a game of %s", minTeams, key),
			Status: bgerr.StatusTooFewTeams,
		}
	} else if len(options.Teams) > maxTeams {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("at most %d teams allowed to create a game of %s", maxTeams, key),
			Status: bgerr.StatusTooManyTeams,
		}
	}
	return &Connect4{
		state:   newState(options.Teams),
		actions: make([]*bg.BoardGameAction, 0),
	}, nil
}

func (c *Connect4) Do(action bg.BoardGameAction) error {
	switch action.ActionType {
	case ActionPlaceDisk:
		var details PlaceDiskActionDetails
		if err := mapstructure.Decode(action.MoreDetails, &details); err != nil {
			return &bgerr.Error{
				Err:    err,
				Status: bgerr.StatusInvalidActionDetails,
			}
		}
		if err := c.state.PlaceDisk(action.Team, details.Column); err != nil {
			return err
		}
		c.actions = append(c.actions, &action)
	case bg.ActionReset:
		c.state = newState(c.state.teams)
		c.actions = make([]*bg.BoardGameAction, 0)
	case bg.ActionUndo:
		if len(c.actions) > 0 {
			undo, _ := NewConnect4(bg.BoardGameOptions{Teams: c.state.teams})
			for _, a := range c.actions[:len(c.actions)-1] {
				if err := undo.Do(*a); err != nil {
					return err
				}
			}
			c.state = undo.state
			c.actions = undo.actions
		} else {
			return &bgerr.Error{
				Err:    fmt.Errorf("no actions to undo"),
				Status: bgerr.StatusInvalidAction,
			}
		}
	default:
		return &bgerr.Error{
			Err:    fmt.Errorf("cannot process action type %s", action.ActionType),
			Status: bgerr.StatusUnknownActionType,
		}
	}
	return nil
}

func (c *Connect4) GetSnapshot(team ...string) (*bg.BoardGameSnapshot, error) {
	if len(team) > 1 {
		return nil, &bgerr.Error{
			Err:    fmt.Errorf("get snapshot requires zero or one team"),
			Status: bgerr.StatusTooManyTeams,
		}
	}
	return &bg.BoardGameSnapshot{
		Turn:    c.state.turn,
		Teams:   c.state.teams,
		Winners: c.state.winners,
		MoreData: Connect4SnapshotDetails{
			Board: c.state.board.board,
		},
		Actions: c.actions,
	}, nil
}

func (c *Connect4) GetSeed() int64 {
	return 0
}
