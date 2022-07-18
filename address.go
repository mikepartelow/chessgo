package chessgo

import "fmt"

type Address string

func NewAddress(file, rank byte) Address {
	return Address(fmt.Sprintf("%c%c", file, rank))
}

func (a Address) File() byte {
	return a[0]
}

func (a Address) Rank() byte {
	return a[1]
}

func (a Address) Plus(incX, incY int8) Address {
	bytePlus := func(b byte, n int8) byte {
		return byte(int8(b) + n)
	}

	return Address(fmt.Sprintf("%c%c", bytePlus(a[0], incX), bytePlus(a[1], incY)))
}
