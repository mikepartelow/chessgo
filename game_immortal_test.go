package chessgo_test

import (
	"fmt"
	"mp/chessgo"
	"testing"
)

// todo: move integration tests to own directory

// https://github.com/mikepartelow/chesspy/blob/main/app/tests/games/immortal.txt

func Test(t *testing.T) {
	moves := []struct {
		boardBefore  string
		move         string
		wantBoard    string
		wantCaptured chessgo.Piece
		wantErr      error
	}{
		{
			boardBefore:  "RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr",
			move:         "e4",
			wantBoard:    "RNBQKBNRPPPP PPP            P                   pppppppprnbqkbnr",
			wantCaptured: chessgo.NoPiece,
			wantErr:      nil,
		},
		{
			boardBefore:  "RNBQKBNRPPPP PPP            P                   pppppppprnbqkbnr",
			move:         "e5",
			wantBoard:    "RNBQKBNRPPPP PPP            P       p           pppp ppprnbqkbnr",
			wantCaptured: chessgo.NoPiece,
			wantErr:      nil,
		},
	}

	board := chessgo.NewBoard()
	game := chessgo.Game{Board: board}

	for moveIndex, tC := range moves {
		desc := fmt.Sprintf("move index %d %s", moveIndex, tC.move)
		t.Run(desc, func(t *testing.T) {
			if game.Board.String() != tC.boardBefore {
				t.Fatalf("Unexpected boardBefore. Want %q, got %q", tC.boardBefore, game.Board.String())
			}

			captured, err := game.Move(tC.move)

			if captured != tC.wantCaptured {
				t.Fatalf("Wanted %q got %q", tC.wantCaptured, captured)
			}

			if err != tC.wantErr {
				t.Fatalf("Wanted %v got %v", tC.wantErr, err)
			}

			if game.Board.String() != tC.wantBoard {
				t.Fatalf("Want Board %q, got %q", tC.wantBoard, game.Board.String())
			}
		})
	}
}
