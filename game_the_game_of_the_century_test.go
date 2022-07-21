package chessgo_test

import "mp/chessgo"

// https://www.chessgames.com/perl/chessgame?gid=1008361

var theGameOftheCenturyMoveExpectations = []moveExpectation{
	{
		move:      "Nf3",
		wantBoard: "RNBQKB RPPPPPPPP     N                          pppppppprnbqkbnr",
	},
	{
		move:      "Nf6",
		wantBoard: "RNBQKB RPPPPPPPP     N                       n  pppppppprnbqkb r",
	},
	{
		move:      "c4",
		wantBoard: "RNBQKB RPP PPPPP     N    P                  n  pppppppprnbqkb r",
	},
	{
		move:      "g6",
		wantBoard: "RNBQKB RPP PPPPP     N    P                  np pppppp prnbqkb r",
	},
	{
		move:      "Nc3",
		wantBoard: "R BQKB RPP PPPPP  N  N    P                  np pppppp prnbqkb r",
	},
	{
		move:      "Bg7",
		wantBoard: "R BQKB RPP PPPPP  N  N    P                  np ppppppbprnbqk  r",
	},
	{
		move:      "d4",
		wantBoard: "R BQKB RPP  PPPP  N  N    PP                 np ppppppbprnbqk  r",
	},
	{
		move:      "O-O",
		wantBoard: "R BQKB RPP  PPPP  N  N    PP                 np ppppppbprnbq rk ",
	},
	{
		move:      "Bf4",
		wantBoard: "R  QKB RPP  PPPP  N  N    PP B               np ppppppbprnbq rk ",
	},
	{
		move:      "d5",
		wantBoard: "R  QKB RPP  PPPP  N  N    PP B     p         np ppp ppbprnbq rk ",
	},
	{
		move:      "Qb3",
		wantBoard: "R   KB RPP  PPPP QN  N    PP B     p         np ppp ppbprnbq rk ",
	},
	{
		move:         "dxc4",
		wantBoard:    "R   KB RPP  PPPP QN  N    pP B               np ppp ppbprnbq rk ",
		wantCaptured: chessgo.WhitePawn{},
	},
	{
		move:         "Qxc4",
		wantBoard:    "R   KB RPP  PPPP  N  N    QP B               np ppp ppbprnbq rk ",
		wantCaptured: chessgo.BlackPawn{},
	},
	{
		move:      "c6",
		wantBoard: "R   KB RPP  PPPP  N  N    QP B            p  np pp  ppbprnbq rk ",
	},
	{
		move:      "e4",
		wantBoard: "R   KB RPP   PPP  N  N    QPPB            p  np pp  ppbprnbq rk ",
	},
	// {
	// 	move:      "Nbd7",
	// 	wantBoard: "R   KB RPP   PPP  N  N    QPPB            p  np pp nppbpr bq rk ",
	// },
}
