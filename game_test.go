package chessgo_test

import (
	"mp/chessgo"
	"testing"
)

func TestGame(t *testing.T) {
	g := chessgo.NewGame()
	assertSquare(t, g.Board(), "e2", 'P')

	// g.Move("e4")
	// assertSquare(t, g.Board().Square("e4"), 'P')
}
