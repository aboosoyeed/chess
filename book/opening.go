package book

import (
	"errors"
	"github.com/andrewbackes/chess/game"
)

// Opening is an  opening to a chess game.
type Opening []Entry

// Apply makes the moves in the opening on the game.
func Apply(opening Opening, g *game.Game) error {
	for _, entry := range opening {
		move, err := g.Position.ParseMove(entry.Move)
		if err != nil {
			return err
		}
		if g.MakeMove(move) != game.InProgress {
			return errors.New("game ended")
		}
	}
	return nil
}

// RandomOpening picks an opening from the book at random.
func (b *Book) RandomOpening(halfmoves int) (Opening, error) {

	return Opening{}, nil
}