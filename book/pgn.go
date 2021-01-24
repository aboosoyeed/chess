package book

import (
	"errors"
	"github.com/aboosoyeed/chess/game"
	"github.com/aboosoyeed/chess/pgn"
	"github.com/aboosoyeed/chess/position"
	"github.com/aboosoyeed/chess/position/move"
)

// FromPGN creates an opening book from a PGN. 'depth' is the number of plies
// to include in the opening book. Variations that arent deep enough aren't
// included. Meaning, if you specify depth 14 but all of the games in your
// pgn only go to depth 10, then your book will be empty.
func FromPGN(pgns []*pgn.PGN, depth int) (*Book, error) {
	if len(pgns) == 0 {
		return nil, errors.New("no games in pgn")
	}
	book := New()
	for _, pgn := range pgns {
		// skip games where we don't know the opening moves
		if pgn.Tags["FEN"] == "" || pgn.Tags["FEN"] == "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1" {
			book.addPgn(pgn, depth)
		}
	}
	return book, nil
}

func (b *Book) addPgn(pgn *pgn.PGN, depth int) {
	g := game.New()
	type pair struct {
		key  position.Hash
		move move.Move
	}
	var staged []pair
	for d, m := range pgn.Moves {
		if d >= depth {
			for _, s := range staged {
				b.addMove(s.key, s.move)
			}
			return
		}
		if mv, err := g.Position().ParseMove(m); err == nil {
			key := g.Position().Polyglot()
			staged = append(staged, pair{key, mv})
			status, _ := g.MakeMove(mv)
			if status != game.InProgress {
				return
			}
		} else {
			return
		}
	}
}

func (b *Book) addMove(key position.Hash, m move.Move) {
	ml := b.Positions[key]
	for i := range ml {
		if ml[i].Move == m {
			ml[i].Weight++
			b.Positions[key] = ml
			return
		}
	}
	b.Positions[key] = append(b.Positions[key], Entry{Move: m, Weight: 1})
}
