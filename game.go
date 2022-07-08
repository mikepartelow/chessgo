package chessgo

type Game struct {
	Board Board
}

func (g *Game) Move(move string) {
	g.Board.SetSquare("a1", ' ')
	g.Board.SetSquare(move, 'P')
}
