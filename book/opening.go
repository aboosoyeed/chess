package book

import (
	"errors"
	"github.com/aboosoyeed/chess/game"
)

// Opening is an  opening to a chess game.
type Opening []Entry

// Apply makes the moves in the opening on the game.
func Apply(opening Opening, g *game.Game) error {
	for _, entry := range opening {
		if s, e := g.MakeMove(entry.Move); e != nil {
			return e
		} else if s != game.InProgress {
			return errors.New("game ended")
		}
	}
	return nil
}

// RandomOpening picks an opening from the book at random.
func (b *Book) RandomOpening(halfmoves int) (Opening, error) {

	return Opening{}, nil
}
