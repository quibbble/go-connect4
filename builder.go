package go_connect4

import (
	"fmt"
	bg "github.com/quibbble/go-boardgame"
	"strconv"
	"strings"
)

const key = "connect4"

type Builder struct{}

func (b *Builder) Create(options *bg.BoardGameOptions) (bg.BoardGame, error) {
	return NewConnect4(options)
}

func (b *Builder) Key() string {
	return key
}

func (b *Builder) CreateAdvanced(options *bg.BoardGameOptions) (bg.AdvancedBoardGame, error) {
	return NewConnect4(options)
}

func (b *Builder) Load(teams []string, notation string) (bg.AdvancedBoardGame, error) {
	// split into four - number teams:seed:options:actions
	splitOne := strings.Split(notation, ":")
	if len(splitOne) != 4 {
		return nil, loadFailure(fmt.Errorf("got %d but wanted %d fields in when decoding", len(splitOne), 4))
	}
	numberTeams, err := strconv.Atoi(splitOne[0])
	if err != nil {
		return nil, loadFailure(err)
	}
	if len(teams) != numberTeams {
		return nil, loadFailure(fmt.Errorf("length of teams %d does not match notation %d", len(teams), numberTeams))
	}
	seed, err := strconv.Atoi(splitOne[1])
	if err != nil {
		return nil, loadFailure(err)
	}
	game, err := NewConnect4(&bg.BoardGameOptions{
		Teams: teams,
		Seed:  int64(seed),
	})
	if err != nil {
		return nil, loadFailure(err)
	}
	// split actions - action;action;action;...
	splitTwo := strings.Split(splitOne[3], ";")
	for _, action := range splitTwo {
		if action == "" {
			break
		}
		// split first two fields of action - team,action type,details,details,...
		splitThree := strings.SplitN(action, ",", 3)
		if len(splitThree) < 2 {
			return nil, loadFailure(fmt.Errorf("got %d but wanted at least %d fields when decoding action", len(splitThree), 2))
		}
		teamIndex, err := strconv.Atoi(splitThree[0])
		if err != nil {
			return nil, loadFailure(err)
		}
		if teamIndex < 0 || teamIndex >= len(teams) {
			return nil, loadFailure(fmt.Errorf("got %d but wanted a team index less than %d when decoding action", teamIndex, len(teams)))
		}
		team := teams[teamIndex]
		actionType := notationIntToAction[splitThree[1]]
		var details interface{}
		if len(splitThree) > 2 {
			switch actionType {
			case ActionPlaceDisk:
				result, err := decodeNotationPlaceDiskActionDetails(splitThree[2])
				if err != nil {
					return nil, err
				}
				details = result
			default:
				return nil, loadFailure(fmt.Errorf("got unneeded action details when decoding action type %s", splitThree[1]))
			}
		}
		if err := game.Do(&bg.BoardGameAction{Team: team, ActionType: actionType, MoreDetails: details}); err != nil {
			return nil, err
		}
	}
	return game, nil
}
