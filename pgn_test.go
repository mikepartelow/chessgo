package chessgo_test

import (
	"mp/chessgo"
	"os"
	"reflect"
	"strings"
	"testing"
)

// [Event "F/S Return Match"]
// [Site "Belgrade, Serbia JUG"]
// [Date "1992.11.04"]
// [Round "29"]
// [White "Fischer, Robert J."]
// [Black "Spassky, Boris V."]
// [Result "1/2-1/2"]

// 1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}
// 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7
// 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5
// Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6
// 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5
// hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5
// 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6
// Nf2 42. g4 Bd3 43. Re6 1/2-1/2

func TestPgn(t *testing.T) {
	t.Run("it parses a tag", func(t *testing.T) {
		want := []chessgo.PGNTag{
			{"Event", "F/S Return Match"},
		}
		text := strings.NewReader(`[Event "F/S Return Match"]`)
		got, _, _ := chessgo.ParsePGN(text)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("it parses several tags", func(t *testing.T) {
		want := []chessgo.PGNTag{
			{"Site", "Belgrade, Serbia JUG"},
			{"Date", "1992.11.04"},
			{"Round", "29"},
		}

		text := strings.NewReader(`[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]`)
		got, _, _ := chessgo.ParsePGN(text)
		assertTags(t, got, want)

	})

	t.Run("it parses movetext", func(t *testing.T) {
		cases := []struct {
			pgn       string
			wantMoves []string
		}{
			{"1. e4 e5", []string{"e4", "e5"}},
			{"1. e4", []string{"e4"}},
			{"1. e4 e5 2. Nf3 Nc6", []string{"e4", "e5", "Nf3", "Nc6"}},
			{"1. e4 e5 {This is a comment.} 2. Ba4 Nf6", []string{"e4", "e5", "Ba4", "Nf6"}},
			{"1. e4 e5 2. Ba4 {This is a comment.} 2... Nf6", []string{"e4", "e5", "Ba4", "Nf6"}},
			{"1. e4 e5 2. Ba4 (This is a comment.) 2... Nf6", []string{"e4", "e5", "Ba4", "Nf6"}},
		}

		for _, tC := range cases {
			t.Run(tC.pgn, func(t *testing.T) {
				text := strings.NewReader(tC.pgn)
				_, got, err := chessgo.ParsePGN(text)

				assertNoError(t, err)
				assertMoves(t, got, tC.wantMoves)
			})

		}
	})

	t.Run("it parses the immortal game", func(t *testing.T) {
		wantTags := []chessgo.PGNTag{
			{"Event", "Adolf Anderssen - Lionel Kieseritzky, 1851 (The Immortal Game): Adolf Anderssen - Lionel Kieseritzky, 1851 (The Immortal Game)"},
			{"Site", "https://lichess.org/study/agESaWbF/cScPsrC4"},
			{"Result", "*"},
			{"UTCDate", "2018.06.18"},
			{"UTCTime", "11:52:03"},
			{"Variant", "Standard"},
			{"ECO", "C33"},
			{"Opening", "King's Gambit Accepted: Bishop's Gambit, Bryan Countergambit"},
			{"Annotator", "https://lichess.org/@/flohahn22"},
		}

		wantMoves := []string{
			"e4", "e5", "f4", "exf4", "Bc4", "Qh4+", "Kf1", "b5", "Bxb5", "Nf6", "Nf3", "Qh6", "d3", "Nh5", "Nh4", "Qg5", "Nf5", "c6", "g4", "Nf6", "Rg1", "cxb5", "h4", "Qg6", "h5", "Qg5", "Qf3", "Ng8", "Bxf4", "Qf6", "Nc3", "Bc5", "Nd5", "Qxb2", "Bd6", "Bxg1", "e5", "Qxa1+", "Ke2", "Na6", "Nxg7+", "Kd8", "Qf6+", "Nxf6", "Be7#",
		}

		f := mustOpen(t, "games/the_immortal_game.pgn")
		gotTags, gotMoves, err := chessgo.ParsePGN(f)

		assertNoError(t, err)
		assertTags(t, gotTags, wantTags)
		assertMoves(t, gotMoves, wantMoves)
	})

	// t.Run("it parses tags plus movetext", func(t *testing.T) {

	// })

	// t.Run("it returns an error on invalid input", func(t *testing.T) {
	// 	text := strings.NewReader("invalid")
	// 	_, _, err := chessgo.ParsePGN(text)

	// 	if err == nil {
	// 		t.Errorf("got nil, wanted error")
	// 	}
	// })
}

func mustOpen(t *testing.T, path string) *os.File {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Errorf("couldn't open %s: %v", path, err)
	}
	return f
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("expected no error, got %s", got)
	}
}

func assertTags(t *testing.T, got, want []chessgo.PGNTag) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func assertMoves(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

}
