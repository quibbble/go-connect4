package go_connect4

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgerr"
	"strconv"
)

var (
	actionToNotation = map[string]string{ActionPlaceDisk: "p"}
	notationToAction = reverseMap(actionToNotation)
)

func (p *PlaceDiskActionDetails) encode() []string {
	return []string{strconv.Itoa(p.Column)}
}

func decodePlaceDiskActionDetails(notation []string) (*PlaceDiskActionDetails, error) {
	if len(notation) != 1 {
		return nil, loadFailure(fmt.Errorf("invalid place disk notation"))
	}
	column, err := strconv.Atoi(notation[0])
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
