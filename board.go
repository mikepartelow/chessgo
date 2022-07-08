package chessgo

import "fmt"

type Board struct {
	squares []byte
}

func NewBoard() *Board {
	return &Board{
		squares: []byte("RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr"),
	}
}

func (b *Board) InBounds(addr string) bool {
	if len(addr) != 2 {
		return false
	}
	index := b.getIndex(addr)
	return index <= uint8(len(b.squares))
}

func (b *Board) GetSquare(addr string) Piece {
	index := b.getIndex(addr)
	return Piece(b.squares[index])
}

func (b *Board) SetSquare(addr string, piece Piece) {
	if !validPiece(piece) {
		panic(fmt.Sprintf("Invalid piece: %c", piece))
	}
	index := b.getIndex(addr)
	b.squares[index] = byte(piece)
}

// panic on out-of-bounds, like a slice would do
// callers can use Board.InBounds() for error checking
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
