package chessgo

import "fmt"

func NewAddress(file, rank byte) string {
	return fmt.Sprintf("%c%c", file, rank)
}

func AddressFile(addr string) byte {
	return addr[0]
}

func AddressRank(addr string) byte {
	return addr[1]
}

func AddressPlus(addr string, incX, incY int8) string {
	bytePlus := func(b byte, n int8) byte {
		return byte(int8(b) + n)
	}

	return fmt.Sprintf("%c%c", bytePlus(addr[0], incX), bytePlus(addr[1], incY))
}
