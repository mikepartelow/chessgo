package chessgo

type Game struct {
	board *Board
}

func NewGame() *Game {
	return &Game{NewBoard()}
}

func (g *Game) Board() *Board {
	return g.board
}
