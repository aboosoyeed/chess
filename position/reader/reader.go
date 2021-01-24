package reader

import (
	"github.com/aboosoyeed/chess/piece"
	"github.com/aboosoyeed/chess/position/board"
	"github.com/aboosoyeed/chess/position/square"
)

// PositionReader is the minimum set of methods a position representation
// needs to implement in order to encode it.
type PositionReader interface {
	OnSquare(square.Square) piece.Piece
	GetCastlingRights() map[piece.Color]map[board.Side]bool
	GetActiveColor() piece.Color
	GetEnPassant() square.Square
	GetFiftyMoveCount() uint64
	GetMoveNumber() int
}
