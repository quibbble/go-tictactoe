package go_tictactoe

import (
	bg "github.com/quibbble/go-boardgame"
)

// key should be set to the name of the game
const key = "Tic-Tac-Toe"

// Builder implements the bg.BoardGameBuilder interface
type Builder struct{}

// Create creates a new game instance
func (b *Builder) Create(options *bg.BoardGameOptions) (bg.BoardGame, error) {
	return NewTicTacToe(options)
}

func (b *Builder) Info() *bg.BoardGameInfo {
	return &bg.BoardGameInfo{
		GameKey:  b.Key(),
		MinTeams: minTeams,
		MaxTeams: maxTeams,
	}
}

// Key returns the key unique to the game i.e. the name of the game
func (b *Builder) Key() string {
	return key
}
