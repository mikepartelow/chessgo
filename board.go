package chessgo

type Board struct {
	squares []byte
}

func NewBoard() *Board {
	return &Board{
		squares: []byte("RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr"),
	}
}

func (b *Board) GetSquare(addr string) Piece {
	index := b.getIndex(addr)
	return Piece(b.squares[index])
}

func (b *Board) SetSquare(addr string, piece Piece) {
	index := b.getIndex(addr)
	b.squares[index] = byte(piece)
}

func (b *Board) getIndex(addr string) uint8 {
	file := 7 - (uint8('h') - addr[0])
	rank := 7 - (uint8('8') - uint8(addr[1]))

	return file + rank*8
}

func (b *Board) Move(fromAddr, toAddr string) Piece {
	replaced := b.GetSquare(toAddr)
	b.SetSquare(toAddr, b.GetSquare(fromAddr))
	b.SetSquare(fromAddr, NoPiece)

	return replaced
}
