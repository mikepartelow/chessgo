package chessgo_test

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
}
