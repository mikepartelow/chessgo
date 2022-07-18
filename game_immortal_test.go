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
		wantCheck    bool
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
		{
			move:      "Bc4",
			wantBoard: "RNBQK NRPPPP  PP          B Pp                  pppp ppprnbqkbnr",
		},
		{
			move:      "Qh4+",
			wantBoard: "RNBQK NRPPPP  PP          B Pp q                pppp ppprnb kbnr",
			wantCheck: true,
		},
		{
			move:      "Kf1",
			wantBoard: "RNBQ KNRPPPP  PP          B Pp q                pppp ppprnb kbnr",
		},
		{
			move:      "b5",
			wantBoard: "RNBQ KNRPPPP  PP          B Pp q p              p pp ppprnb kbnr",
		},
		{
			move:      "Bxb5",
			wantBoard: "RNBQ KNRPPPP  PP            Pp q B              p pp ppprnb kbnr",
		},
		{
			move:      "Nf6",
			wantBoard: "RNBQ KNRPPPP  PP            Pp q B           n  p pp ppprnb kb r",
		},
		{
			move:      "Nf3",
			wantBoard: "RNBQ K RPPPP  PP     N      Pp q B           n  p pp ppprnb kb r",
		},
		{
			move:      "Qh6",
			wantBoard: "RNBQ K RPPPP  PP     N      Pp   B           n qp pp ppprnb kb r",
		},
		{
			move:      "d3",
			wantBoard: "RNBQ K RPPP   PP   P N      Pp   B           n qp pp ppprnb kb r",
		},
		{
			move:      "Nh5",
			wantBoard: "RNBQ K RPPP   PP   P N      Pp   B     n       qp pp ppprnb kb r",
		},
		{
			move:      "Nh4",
			wantBoard: "RNBQ K RPPP   PP   P        Pp N B     n       qp pp ppprnb kb r",
		},
		{
			move:      "Qg5",
			wantBoard: "RNBQ K RPPP   PP   P        Pp N B    qn        p pp ppprnb kb r",
		},
		{
			move:      "Nf5",
			wantBoard: "RNBQ K RPPP   PP   P        Pp   B   Nqn        p pp ppprnb kb r",
		},
		{
			move:      "c6",
			wantBoard: "RNBQ K RPPP   PP   P        Pp   B   Nqn  p     p  p ppprnb kb r",
		},
		{
			move:      "g4",
			wantBoard: "RNBQ K RPPP    P   P        PpP  B   Nqn  p     p  p ppprnb kb r",
		},
		{
			move:      "Nf6",
			wantBoard: "RNBQ K RPPP    P   P        PpP  B   Nq   p  n  p  p ppprnb kb r",
		},
		{
			move:      "Rg1",
			wantBoard: "RNBQ KR PPP    P   P        PpP  B   Nq   p  n  p  p ppprnb kb r",
		},
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

			if game.Board.InCheck(game.Turn) != tC.wantCheck {
				t.Fatalf("Wanted Check but Board reports no check.")
			}
		})
	}
}
