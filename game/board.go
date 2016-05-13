// Package game plays chess games.
package game

import (
	"fmt"
	"strconv"
	"strings"
)

// Board is a representation of a chess Board.
type Board struct {
	// bitBoard has one bitBoard per player per color.
	bitBoard [2][6]uint64 //[player][piece]
}

func NewBoard() Board {
	b := Board{bitBoard: [2][6]uint64{}}
	b.Reset()
	return b
}

// String puts the Board into a pretty print-able format.
func (b Board) String() (str string) {
	str += "+---+---+---+---+---+---+---+---+\n"
	for i := 1; i <= 64; i++ {
		square := Square(64 - i)
		str += "|"
		noPiece := true
		for c := range b.bitBoard {
			for j := range b.bitBoard[c] {
				if ((1 << square) & b.bitBoard[c][j]) != 0 {
					str += fmt.Sprint(" ", b.OnSquare(square), " ")
					noPiece = false
				}
			}
		}
		if noPiece {
			str += "   "
		}
		if square%8 == 0 {
			str += "|\n"
			str += "+---+---+---+---+---+---+---+---+"
			if square < LastSquare {
				str += "\n"
			}
		}
	}
	return
}

// Clear empties the Board.
func (b *Board) Clear() {
	b.bitBoard = [2][6]uint64{}
}

// Reset puts the pieces in the new game position.
func (b *Board) Reset() {
	// puts the pieces in their starting/newgame positions
	for color := uint(0); color < 2; color = color + 1 {
		//Pawns first:
		b.bitBoard[color][Pawn] = 255 << (8 + (color * 8 * 5))
		//Then the rest of the pieces:
		b.bitBoard[color][Knight] = (1 << (B1 + (color * 8 * 7))) ^ (1 << (G1 + (color * 8 * 7)))
		b.bitBoard[color][Bishop] = (1 << (C1 + (color * 8 * 7))) ^ (1 << (F1 + (color * 8 * 7)))
		b.bitBoard[color][Rook] = (1 << (A1 + (color * 8 * 7))) ^ (1 << (H1 + (color * 8 * 7)))
		b.bitBoard[color][Queen] = (1 << (D1 + (color * 8 * 7)))
		b.bitBoard[color][King] = (1 << (E1 + (color * 8 * 7)))
	}
}

// OnSquare returns the piece that is on the specified square.
func (b *Board) OnSquare(s Square) Piece {
	for c := White; c <= Black; c++ {
		for p := Pawn; p <= King; p++ {
			if (b.bitBoard[c][p] & (1 << s)) != 0 {
				return NewPiece(c, p)
			}
		}
	}
	return NewPiece(Neither, None)
}

// Occupied returns a bitBoard with all of the specified colors pieces.
func (b *Board) occupied(c Color) uint64 {
	var mask uint64
	for p := Pawn; p <= King; p++ {
		if c == Both {
			mask |= b.bitBoard[White][p] | b.bitBoard[Black][p]
		} else {
			mask |= b.bitBoard[c][p]
		}
	}
	return mask
}

// MakeMove attempts to make the given move no matter legality.
func (b *Board) MakeMove(m Move) {
	from, to := getSquares(m)
	movingPiece := b.OnSquare(from)
	capturedPiece := b.OnSquare(to)

	// Remove captured piece:
	if capturedPiece.Type != None {
		b.bitBoard[capturedPiece.Color][capturedPiece.Type] ^= (1 << to)
	}

	// Move piece:
	b.bitBoard[movingPiece.Color][movingPiece.Type] ^= ((1 << from) | (1 << to))

	// Castle:
	if movingPiece.Type == King {
		if from == Square(E1+56*uint8(movingPiece.Color)) && (to == Square(G1+56*uint8(movingPiece.Color))) {
			b.bitBoard[movingPiece.Color][Rook] ^= (1 << (H1 + 56*uint8(movingPiece.Color))) | (1 << (F1 + 56*uint8(movingPiece.Color)))
		} else if from == Square(E1+56*uint8(movingPiece.Color)) && to == Square(C1+56*uint8(movingPiece.Color)) {
			b.bitBoard[movingPiece.Color][Rook] ^= (1 << (A1 + 56*uint8(movingPiece.Color))) | (1 << (D1 + 56*uint8(movingPiece.Color)))
		}
	}

	if movingPiece.Type == Pawn {
		// Handle en Passant capture:
		// capturedPiece just means the piece on the destination square
		if (int(to)-int(from))%8 != 0 && capturedPiece.Type == None {
			if movingPiece.Color == White {
				b.bitBoard[Black][Pawn] ^= (1 << (to - 8))
			} else if movingPiece.Color == Black {
				b.bitBoard[White][Pawn] ^= (1 << (to + 8))
			}
		}
		// Handle Promotions:
		promotesTo := promotedPiece(m)
		if promotesTo != NoPiece {
			b.bitBoard[movingPiece.Color][movingPiece.Type] ^= (1 << to) // remove Pawn
			b.bitBoard[movingPiece.Color][promotesTo] ^= (1 << to)       // add promoted piece
		}
	}

}

// parseBoard parses the board passed via FEN and returns a board object.
func parseBoard(position string) *Board {
	b := NewBoard()
	b.Clear()
	// remove the /'s and replace the numbers with that many spaces
	// so that there is a 1-1 mapping from bytes to squares.
	parsedBoard := strings.Replace(position, "/", "", 9)
	for i := 1; i < 9; i++ {
		parsedBoard = strings.Replace(parsedBoard, strconv.Itoa(i), strings.Repeat(" ", i), -1)
	}
	piece := map[string]PieceType{
		"P": Pawn, "p": Pawn,
		"N": Knight, "n": Knight,
		"B": Bishop, "b": Bishop,
		"R": Rook, "r": Rook,
		"Q": Queen, "q": Queen,
		"K": King, "k": King}
	color := map[string]Color{
		"P": White, "p": Black,
		"N": White, "n": Black,
		"B": White, "b": Black,
		"R": White, "r": Black,
		"Q": White, "q": Black,
		"K": White, "k": Black}
	// adjust the bitboards:
	for pos := 0; pos < len(parsedBoard); pos++ {
		k := parsedBoard[pos:(pos + 1)]
		if _, ok := piece[k]; ok {
			b.bitBoard[color[k]][piece[k]] |= (1 << uint(63-pos))
		}
	}
	return &b
}

// Put places a piece on the square.
func (b *Board) Put(p Piece, s Square) {
	b.bitBoard[p.Color][p.Type] |= (1 << s)
}

func (b *Board) printbitBoards() {
	for c := range b.bitBoard {
		for j := range b.bitBoard[c] {
			fmt.Println(NewPiece(Color(c), PieceType(j)))
			bitprint(b.bitBoard[c][j])
		}
	}
}
