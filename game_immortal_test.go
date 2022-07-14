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
		move         string
		wantBoard    string
		wantCaptured chessgo.Piece
	}{
		{

			move:      "e4",
			wantBoard: "RNBQKBNRPPPP PPP            P                   pppppppprnbqkbnr",
		},
		{
			move:      "e5",
			wantBoard: "RNBQKBNRPPPP PPP            P       p           pppp ppprnbqkbnr",
		},
		{
			move:      "f4",
			wantBoard: "RNBQKBNRPPPP  PP            PP      p           pppp ppprnbqkbnr",
		},
		{
			move:      "exf4",
			wantBoard: "RNBQKBNRPPPP  PP            Pp                  pppp ppprnbqkbnr",
		},
		// {
		// 	move:      "Bc4",
		// 	wantBoard: "RNBQK NRPPPP  PP          B Pp                  pppp ppprnbqkbnr",
		// },
	}

	board := chessgo.NewBoard()
	game := chessgo.Game{Board: board}

	for moveIndex, tC := range moves {
		desc := fmt.Sprintf("move index %d %s", moveIndex, tC.move)
		t.Run(desc, func(t *testing.T) {
			captured, err := game.Move(tC.move)

			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}

			if tC.wantCaptured != chessgo.Piece(0) && captured != tC.wantCaptured {
				t.Fatalf("Wanted %q got %q", tC.wantCaptured, captured)
			}

			if game.Board.String() != tC.wantBoard {
				t.Fatalf("Want Board %q, got %q", tC.wantBoard, game.Board.String())
			}
		})
	}
}
