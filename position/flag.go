package position

import (
	"github.com/aboosoyeed/chess/piece"
	"github.com/aboosoyeed/chess/position/square"
	"strings"
)

const (
	BITS_NORMAL       int = 1
	BITS_CAPTURE      int = 2
	BITS_EP_CAPTURE   int = 3
	BITS_PROMOTION    int = 4
	BITS_KSIDE_CASTLE int = 5
	BITS_QSIDE_CASTLE int = 6
)

func GetFlag(san string, movingPiece piece.Piece, from, to square.Square, p *Position) int {
	if san == "O-O" {
		return BITS_KSIDE_CASTLE
	} else if san == "O-O-O" {
		return BITS_QSIDE_CASTLE
	} else if strings.Contains(san, "=") {
		return BITS_PROMOTION
	} else if strings.Contains(san, "x") {
		if movingPiece.Type == piece.Pawn && isEnPassant(p, to) {
			return BITS_EP_CAPTURE
		}
		return BITS_CAPTURE
	}
	return BITS_NORMAL
}

func isEnPassant(pos *Position, to square.Square) bool {
	p := pos.OnSquare(to)
	return p.Type != piece.Pawn
}
