package chessgo

import (
	"fmt"
)

type StandardBoard struct {
	squares []byte
}

func NewBoard() *StandardBoard {
	return NewBoardFromString("RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr")
}

func NewBoardFromString(from string) *StandardBoard {
	if len(from) != 64 {
		panic(fmt.Sprintf("Standard Board is 8x8, can't initialize with %d bytes given.", len(from)))
	}
	return &StandardBoard{squares: []byte(from)}
}

func (b *StandardBoard) InBounds(addr string) bool {
	if len(addr) != 2 {
		return false
	}
	index := b.getIndex(addr)
	return index < uint8(len(b.squares))
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
	// log.Printf(" getIndex(%s)", addr)
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

func (b *StandardBoard) InCheck(color Color) bool {
	kingAddr := b.findKing(color)

	if b.inCheckHorizontal(kingAddr, Queen(ToggleColor(color))) {
		return true
	}

	if b.inCheckDiagonal(kingAddr, Queen(ToggleColor(color))) {
		return true
	}

	return false
}

func (b *StandardBoard) findKing(color Color) (addr string) {
	wantedKing := King(color)

	for i := uint8(0); i < 8; i++ {
		for j := uint8(0); j < 8; j++ {
			addr = fmt.Sprintf("%c%c", 'a'+i, '1'+j)
			if b.GetSquare(addr) == wantedKing {
				return
			}
		}
	}

	panic(fmt.Sprintf("Missing %s King!", color.String()))
}

func (b *StandardBoard) inCheckHorizontal(kingAddr string, opponent Piece) bool {
	increments := []struct {
		incX int8
		incY int8
	}{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for _, incs := range increments {
		incX, incY := incs.incX, incs.incY
		for i, j := incX, incY; i > -7 && i < 8 && j > -7 && j < 8; i, j = i+incX, j+incY {
			addr := AddressPlus(kingAddr, i, j)
			if !b.InBounds(addr) {
				break
			}
			piece := b.GetSquare(addr)
			if piece == opponent {
				return true
			} else if piece != NoPiece {
				break
			}
		}
	}

	return false
}

func (b *StandardBoard) inCheckDiagonal(kingAddr string, opponent Piece) bool {
	increments := []struct {
		incX int8
		incY int8
	}{
		{-1, -1}, {1, 1}, {1, -1}, {-1, 1},
	}

	for _, incs := range increments {
		incX, incY := incs.incX, incs.incY
		for i, j := incX, incY; i > -7 && i < 8 && j > -7 && j < 8; i, j = i+incX, j+incY {
			addr := AddressPlus(kingAddr, i, j)
			if !b.InBounds(addr) {
				break
			}
			piece := b.GetSquare(addr)
			if piece == opponent {
				return true
			} else if piece != NoPiece {
				break
			}
		}
	}

	return false
}
