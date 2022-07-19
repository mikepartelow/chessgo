package chessgo

type WhitePawn struct{}

func (p WhitePawn) Byte() byte     { return 'P' }
func (p WhitePawn) Color() Color   { return White }
func (p WhitePawn) String() string { return "White Pawn" }
func (p WhitePawn) SourceForDest(dstAddr Address, board Board) Address {
	dstPiece := board.GetSquare(dstAddr)
	isSrc := func(addr Address) bool {
		return board.InBounds(addr) && board.GetSquare(addr) == p
	}

	if dstPiece != nil {
		for _, incs := range []increments{{-1, -1}, {1, -1}} {
			addr := dstAddr.Plus(incs.incX, incs.incY)
			if isSrc(addr) {
				return addr
			}
		}
		return ""
	}

	addr := dstAddr.Plus(0, -1)
	if isSrc(addr) {
		return addr
	}

	// first White pawn move can be 2 squares starting from rank 2
	addr = NewAddress(dstAddr.File(), '2')
	if dstAddr.Rank() == '4' && isSrc(addr) {
		return addr
	}

	return ""
}

type BlackPawn struct{}

func (p BlackPawn) Byte() byte     { return 'p' }
func (p BlackPawn) Color() Color   { return Black }
func (p BlackPawn) String() string { return "Black Pawn" }
func (p BlackPawn) SourceForDest(dstAddr Address, board Board) Address {
	dstPiece := board.GetSquare(dstAddr)
	isSrc := func(addr Address) bool {
		return board.InBounds(addr) && board.GetSquare(addr) == p
	}

	if dstPiece != nil {

		for _, incs := range []increments{{1, 1}, {-1, 1}} {
			addr := dstAddr.Plus(incs.incX, incs.incY)
			if isSrc(addr) {
				return addr
			}
		}
		return ""
	}

	addr := dstAddr.Plus(0, 1)
	if isSrc(addr) {
		return addr
	}

	// first Black pawn move can be 2 squares starting from Board.MaxRank() - 1
	addr = NewAddress(dstAddr.File(), byte(board.MaxRank())-1)
	if dstAddr.Rank() == byte(board.MaxRank()-3) && isSrc(addr) {
		return addr
	}

	return ""
}

func Pawn(color Color) Piece {
	switch color {
	case White:
		return WhitePawn{}
	case Black:
		return BlackPawn{}
	default:
		panic("invalid Color")
	}
}
