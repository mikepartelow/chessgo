package chessgo

import "fmt"

type StandardBoard struct {
	squares []byte
}

func NewBoard() *StandardBoard {
	return &StandardBoard{
		squares: []byte("RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr"),
	}
}

func (b *StandardBoard) InBounds(addr string) bool {
	if len(addr) != 2 {
		return false
	}
	index := b.getIndex(addr)
	return index <= uint8(len(b.squares))
}

func (b *StandardBoard) GetSquare(addr string) Piece {
	index := b.getIndex(addr)
	return Piece(b.squares[index])
}

func (b *StandardBoard) SetSquare(addr string, piece Piece) {
	if !validPiece(piece) {
		panic(fmt.Sprintf("Invalid piece: %c", piece))
	}
	index := b.getIndex(addr)
	b.squares[index] = byte(piece)
}

// no bounds checking: panic on out-of-bounds, like a slice would do
// callers can use Board.InBounds() for error checking
func (b *StandardBoard) getIndex(addr string) uint8 {
	file := 7 - (uint8('h') - addr[0])
	rank := 7 - (uint8('8') - uint8(addr[1]))

	return file + rank*8
}

func (b *StandardBoard) Move(srcAddr, dstAddr string) Piece {
	replaced := b.GetSquare(dstAddr)
	b.SetSquare(dstAddr, b.GetSquare(srcAddr))
	b.SetSquare(srcAddr, NoPiece)

	return replaced
}

func (b *StandardBoard) String() string {
	return string(b.squares)
}

func (b *StandardBoard) MaxFile() rune {
	return 'h'
}

func (b *StandardBoard) MaxRank() rune {
	return '8'
}
