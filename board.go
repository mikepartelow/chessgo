package chess

type Board struct {
	squares [8][8]rune
}

type Square struct {
	X uint8 // File
	Y uint8 // Rank
}

type Piece rune

can all this
- [ ] write ACCEPTANCE TEST 0 : board.Move("e4")
- [ ] write a main.go (in folder, see lgwt)

When you're creating packages, even if they're only internal to your project,
 prefer a top-down consumer driven approach. This will stop you over-imagining
 designs and making abstractions you may not even need and will help ensure
 the tests you write are useful.


func (b *Board) Ranks() uint8 { return 8 } // Rows
func (b *Board) Files() uint8 { return 8 } // Columns

func NewBoard() *Board {
	return &Board{
		[8][8]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{'R', 'N', 'B', ' ', ' ', ' ', ' ', ' '},
		},
	}
}

func (b *Board) InBounds(square Square) bool {
	return square.X < b.Files() && square.Y < b.Ranks()
}

func (b *Board) GetPieceAt(square Square) rune {
	return b.squares[square.Y][square.X]
}
