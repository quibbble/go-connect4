package go_connect4

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	bg "github.com/quibbble/go-boardgame"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"github.com/quibbble/go-boardgame/pkg/bgn"
	"math/rand"
	"strings"
)

const (
	minTeams = 2
	maxTeams = 3
)

type Connect4 struct {
	state   *state
	actions []*bg.BoardGameAction
	seed    int64
}

func NewConnect4(options *bg.BoardGameOptions) (*Connect4, error) {
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
		state:   newState(options.Teams, rand.New(rand.NewSource(options.Seed))),
		actions: make([]*bg.BoardGameAction, 0),
		seed:    options.Seed,
	}, nil
}

func (c *Connect4) Do(action *bg.BoardGameAction) error {
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
		c.actions = append(c.actions, action)
	case bg.ActionSetWinners:
		var details bg.SetWinnersActionDetails
		if err := mapstructure.Decode(action.MoreDetails, &details); err != nil {
			return &bgerr.Error{
				Err:    err,
				Status: bgerr.StatusInvalidActionDetails,
			}
		}
		if err := c.state.SetWinners(details.Winners); err != nil {
			return err
		}
		c.actions = append(c.actions, action)
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
	var targets []*bg.BoardGameAction
	if len(c.state.winners) == 0 && (len(team) == 0 || (len(team) == 1 && team[0] == c.state.turn)) {
		targets = c.state.targets()
	}
	return &bg.BoardGameSnapshot{
		Turn:    c.state.turn,
		Teams:   c.state.teams,
		Winners: c.state.winners,
		MoreData: Connect4SnapshotDetails{
			Board: c.state.board.board,
		},
		Targets: targets,
		Actions: c.actions,
	}, nil
}

func (c *Connect4) GetBGN() *bgn.Game {
	tags := map[string]string{
		"Game":  key,
		"Teams": strings.Join(c.state.teams, ", "),
		"Seed":  fmt.Sprintf("%d", c.seed),
	}
	actions := make([]bgn.Action, 0)
	for _, action := range c.actions {
		bgnAction := bgn.Action{
			TeamIndex: indexOf(c.state.teams, action.Team),
			ActionKey: rune(actionToNotation[action.ActionType][0]),
		}
		switch action.ActionType {
		case ActionPlaceDisk:
			var details PlaceDiskActionDetails
			_ = mapstructure.Decode(action.MoreDetails, &details)
			bgnAction.Details = details.encode()
		case bg.ActionSetWinners:
			var details bg.SetWinnersActionDetails
			_ = mapstructure.Decode(action.MoreDetails, &details)
			bgnAction.Details, _ = details.EncodeBGN(c.state.teams)
		}
		actions = append(actions, bgnAction)
	}
	return &bgn.Game{
		Tags:    tags,
		Actions: actions,
	}
}
