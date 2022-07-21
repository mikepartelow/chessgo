package chessgo_test

import (
	"fmt"
	"mp/chessgo"
	"testing"
)

// https://www.chessgames.com/perl/chessgame?gid=1018910

type moveExpectation struct {
	move         string
	wantBoard    string
	wantCaptured chessgo.Piece
	wantCheck    bool
	wantMate     bool
}

func TestFamousGames(t *testing.T) {
	famousGames := []struct {
		name             string
		moveExpectations []moveExpectation
	}{
		{
			name:             "The Immortal Game",
			moveExpectations: theImmortalGameMoveExpectations,
		},
		{
			name:             "The Game Of The Century",
			moveExpectations: theGameOftheCenturyMoveExpectations,
		},
	}

	for _, famousGame := range famousGames {
		t.Run(famousGame.name, func(t *testing.T) {
			board := chessgo.NewBoard()
			game := &chessgo.Game{Board: board}

			for moveIndex, mE := range famousGame.moveExpectations {
				desc := fmt.Sprintf("move index %d %s", moveIndex, mE.move)
				t.Run(desc, func(t *testing.T) {
					assertMoveExpectation(t, mE, game)
				})
			}
		})
	}
}

func assertCaptured(t *testing.T, got, want chessgo.Piece) {
	t.Helper()

	if got != want {
		t.Fatalf("got %q captured, wanted %q", got, want)
	}
}

func assertBoard(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got Board %q, wanted %q", got, want)
	}
}

func assertCheck(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Fatalf("got check == %v, wanted %v", got, want)
	}
}

func assertMoveExpectation(t *testing.T, me moveExpectation, game *chessgo.Game) {
	t.Helper()

	captured, err := game.Move(me.move)

	assertNoError(t, err)
	assertCaptured(t, captured, me.wantCaptured)
	assertBoard(t, game.Board.String(), me.wantBoard)
	assertCheck(t, game.Board.InCheck(game.Turn), me.wantCheck)

	//todo: check for mate!
	// gotMate := game.Board.Mated(game.Turn)
	// if gotMate != tC.wantMate {
	// 	t.Fatalf("Wanted mate == %v but Board reports mate == %v.", tC.wantMate, gotMate)
	// }
}
