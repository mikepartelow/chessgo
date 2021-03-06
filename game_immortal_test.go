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

var immortalGameMoves = []moveExpectation{
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
		move:         "exf4",
		wantBoard:    "RNBQKBNRPPPP  PP            Pp                  pppp ppprnbqkbnr",
		wantCaptured: chessgo.WhitePawn{},
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
		move:         "Bxb5",
		wantBoard:    "RNBQ KNRPPPP  PP            Pp q B              p pp ppprnb kbnr",
		wantCaptured: chessgo.BlackPawn{},
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
	{
		move:         "cxb5",
		wantBoard:    "RNBQ KR PPP    P   P        PpP  p   Nq      n  p  p ppprnb kb r",
		wantCaptured: chessgo.WhiteBishop{},
	},
	{
		move:      "h4",
		wantBoard: "RNBQ KR PPP        P        PpPP p   Nq      n  p  p ppprnb kb r",
	},
	{
		move:      "Qg6",
		wantBoard: "RNBQ KR PPP        P        PpPP p   N       nq p  p ppprnb kb r",
	},
	{
		move:      "h5",
		wantBoard: "RNBQ KR PPP        P        PpP  p   N P     nq p  p ppprnb kb r",
	},
	{
		move:      "Qg5",
		wantBoard: "RNBQ KR PPP        P        PpP  p   NqP     n  p  p ppprnb kb r",
	},
	{
		move:      "Qf3",
		wantBoard: "RNB  KR PPP        P Q      PpP  p   NqP     n  p  p ppprnb kb r",
	},
	{
		move:      "Ng8",
		wantBoard: "RNB  KR PPP        P Q      PpP  p   NqP        p  p ppprnb kbnr",
	},
	{
		move:         "Bxf4",
		wantBoard:    "RN   KR PPP        P Q      PBP  p   NqP        p  p ppprnb kbnr",
		wantCaptured: chessgo.BlackPawn{},
	},
	{
		move:      "Qf6",
		wantBoard: "RN   KR PPP        P Q      PBP  p   N P     q  p  p ppprnb kbnr",
	},
	{
		move:      "Nc3",
		wantBoard: "R    KR PPP       NP Q      PBP  p   N P     q  p  p ppprnb kbnr",
	},
	{
		move:      "Bc5",
		wantBoard: "R    KR PPP       NP Q      PBP  pb  N P     q  p  p ppprnb k nr",
	},
	{
		move:      "Nd5",
		wantBoard: "R    KR PPP        P Q      PBP  pbN N P     q  p  p ppprnb k nr",
	},
	{
		move:         "Qxb2",
		wantBoard:    "R    KR PqP        P Q      PBP  pbN N P        p  p ppprnb k nr",
		wantCaptured: chessgo.WhitePawn{},
	},
	{
		move:      "Bd6",
		wantBoard: "R    KR PqP        P Q      P P  pbN N P   B    p  p ppprnb k nr",
	},
	{
		move:         "Bxg1",
		wantBoard:    "R    Kb PqP        P Q      P P  p N N P   B    p  p ppprnb k nr",
		wantCaptured: chessgo.WhiteRook{},
	},
	{
		move:      "e5",
		wantBoard: "R    Kb PqP        P Q        P  p NPN P   B    p  p ppprnb k nr",
	},
	{
		move:         "Qxa1+",
		wantBoard:    "q    Kb P P        P Q        P  p NPN P   B    p  p ppprnb k nr",
		wantCheck:    true,
		wantCaptured: chessgo.WhiteRook{},
	},
	{
		move:      "Ke2",
		wantBoard: "q     b P P K      P Q        P  p NPN P   B    p  p ppprnb k nr",
	},
	{
		move:      "Na6",
		wantBoard: "q     b P P K      P Q        P  p NPN Pn  B    p  p pppr b k nr",
	},
	{
		move:         "Nxg7+",
		wantBoard:    "q     b P P K      P Q        P  p NP  Pn  B    p  p pNpr b k nr",
		wantCheck:    true,
		wantCaptured: chessgo.BlackPawn{},
	},
	{
		move:      "Kd8",
		wantBoard: "q     b P P K      P Q        P  p NP  Pn  B    p  p pNpr bk  nr",
	},
	{
		move:      "Qf6+",
		wantBoard: "q     b P P K      P          P  p NP  Pn  B Q  p  p pNpr bk  nr",
		wantCheck: true,
	},
	{
		move:         "Nxf6",
		wantBoard:    "q     b P P K      P          P  p NP  Pn  B n  p  p pNpr bk   r",
		wantCaptured: chessgo.WhiteQueen{},
	},
	{
		move:      "Be7#",
		wantBoard: "q     b P P K      P          P  p NP  Pn    n  p  pBpNpr bk   r",
		wantCheck: true,
		wantMate:  true,
	},
}

func TestTheImmortalGame(t *testing.T) {
	t.Run("from memory", func(t *testing.T) {
		board := chessgo.NewBoard()
		game := &chessgo.Game{Board: board}

		for moveIndex, tC := range immortalGameMoves {
			desc := fmt.Sprintf("move index %d %s", moveIndex, tC.move)
			t.Run(desc, func(t *testing.T) {
				assertMoveExpectation(t, tC, game)
			})
		}
	})
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
