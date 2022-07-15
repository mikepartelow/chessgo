package chessgo

import "fmt"

func AddressPlus(addr string, incX, incY int8) string {
	return fmt.Sprintf("%c%c", bytePlus(addr[0], incX), bytePlus(addr[1], incY))
}

func bytePlus(b byte, n int8) byte {
	return byte(int8(b) + n)
}
