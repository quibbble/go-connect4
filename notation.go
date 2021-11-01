package go_connect4

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"strconv"
)

// Notation - "'number of teams':'seed':'MoreOptions':'team index','action type number','details','details';..."

var (
	notationActionToInt = map[string]int{ActionPlaceDisk: 0}
	notationIntToAction = map[string]string{"0": ActionPlaceDisk}
)

func (p *PlaceDiskActionDetails) encode() string {
	return fmt.Sprintf("%d", p.Column)
}

func decodeNotationPlaceDiskActionDetails(notation string) (*PlaceDiskActionDetails, error) {
	column, err := strconv.Atoi(notation)
	if err != nil {
		return nil, loadFailure(err)
	}
	return &PlaceDiskActionDetails{
		Column: column,
	}, nil
}

func loadFailure(err error) error {
	return &bgerr.Error{
		Err:    err,
		Status: bgerr.StatusGameLoadFailure,
	}
}
