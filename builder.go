package go_connect4

import (
	bg "github.com/quibbble/go-boardgame"
)

const key = "connect4"

type Builder struct{}

func (b *Builder) Create(options bg.BoardGameOptions) (bg.BoardGame, error) {
	return NewConnect4(options)
}

func (b *Builder) Key() string {
	return key
}
