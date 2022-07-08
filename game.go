package chessgo

type Game struct {
	Board Board
}

func (g *Game) SourceForDest(dstAddr string) string {
	return "a1"
}
