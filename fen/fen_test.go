package fen

import (
	"github.com/aboosoyeed/chess/game"
	"github.com/aboosoyeed/chess/piece"
	"github.com/aboosoyeed/chess/position"
	"github.com/aboosoyeed/chess/position/board"
	"github.com/aboosoyeed/chess/position/move"
	"github.com/aboosoyeed/chess/position/square"
	"strings"
	"testing"
)

func TestLoadRootPos(t *testing.T) {
	root := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	_, err := Decode(root)
	if err != nil {
		t.Fail()
	}
}

// integration test
func TestFENEncodeDecode(t *testing.T) {
	fens := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		"4k3/8/8/8/8/8/4P3/4K3 w - - 5 39",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
	}
	for _, fen := range fens {
		t.Log("In:  ", fen)
		g, err := Decode(fen)
		out, _ := Encode(g)
		t.Log("Out: ", out)
		if err != nil || fen != out {
			t.Error("Do not match.")
		}
	}

}

func TestFENenPassant(t *testing.T) {
	fen := "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"
	g, _ := Decode(fen)
	if g.EnPassant != square.C6 {
		t.Fail()
	}
}

func TestFENCastlingRights(t *testing.T) {
	fen := "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"
	p, _ := Decode(fen)
	if !p.CastlingRights[piece.White][board.ShortSide] || !p.CastlingRights[piece.Black][board.ShortSide] ||
		!p.CastlingRights[piece.White][board.LongSide] || !p.CastlingRights[piece.Black][board.LongSide] {

		t.Fail()
	}
}

func TestFENMarshalRoot(t *testing.T) {
	fen, _ := Encode(position.New())
	root := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	if fen != root {
		t.Fail()
	}
}

func TestFENWhitesMove(t *testing.T) {
	fen, _ := Encode(position.New())
	player := strings.Split(fen, " ")[1]
	if player != "w" {
		t.Fail()
	}
}

func TestFENBlacksMove(t *testing.T) {
	g := game.New()
	g.MakeMove(move.Parse("e2e4"))
	fen, _ := Encode(g.Position())
	player := strings.Split(fen, " ")[1]
	if player != "b" {
		t.Fail()
	}
}
