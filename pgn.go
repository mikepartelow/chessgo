package chessgo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type PGNTag struct {
	Tag   string
	Value string
}

type parserState struct {
	inComment bool
	turn      Color
	moveIdx   int
}

func ParsePGN(in io.Reader) ([]PGNTag, []string, error) {
	var tags []PGNTag
	var moves []string

	scanner := bufio.NewScanner(in)
	state := parserState{moveIdx: 1}

	for scanner.Scan() {
		line := scanner.Text()
		if tag := parseTag(line); tag != nil {
			tags = append(tags, *tag)
		} else {
			newMoves, err := parseMoves(line, &state)
			if err != nil {
				return nil, nil, fmt.Errorf("problem parsing moves: %v", err)
			}
			moves = append(moves, newMoves...)
		}
	}

	if len(tags) > 0 || len(moves) > 0 {
		return tags, moves, nil
	}

	return nil, nil, fmt.Errorf("invalid input")
}

func parseTag(line string) *PGNTag {
	// [Site "Belgrade, Serbia JUG"]
	// [Date "1992.11.04"]
	// [Round "29"]
	// [White "Fischer, Robert J."]
	// [Black "Spassky, Boris V."]
	// [Result "1/2-1/2"]

	if isBracketedWith("[", "]", line) {
		parts := strings.SplitN(line[1:len(line)-1], " ", 2)

		if len(parts) == 2 && isBracketedWith(`"`, `"`, parts[1]) {
			tag := parts[0]
			value := parts[1][1 : len(parts[1])-1]
			return &PGNTag{tag, value}
		}
	}
	return nil
}

func isBracketedWith(leftBracket, rightBracket, text string) bool {
	return strings.HasPrefix(text, leftBracket) && strings.HasSuffix(text, rightBracket)
}

func parseMoves(line string, state *parserState) ([]string, error) {
	// 1. e4 e5 2. Nf3 Nc6 3. Bb5 a6 {This opening is called the Ruy Lopez.}
	// 4. Ba4 Nf6 5. O-O Be7 6. Re1 b5 7. Bb3 d6 8. c3 O-O 9. h3 Nb8 10. d4 Nbd7
	// 11. c4 c6 12. cxb5 axb5 13. Nc3 Bb7 14. Bg5 b4 15. Nb1 h6 16. Bh4 c5 17. dxe5
	// Nxe4 18. Bxe7 Qxe7 19. exd6 Qf6 20. Nbd2 Nxd6 21. Nc4 Nxc4 22. Bxc4 Nb6
	// 23. Ne5 Rae8 24. Bxf7+ Rxf7 25. Nxf7 Rxe1+ 26. Qxe1 Kxf7 27. Qe3 Qg5 28. Qxg5
	// hxg5 29. b3 Ke6 30. a3 Kd6 31. axb4 cxb4 32. Ra5 Nd5 33. f3 Bc8 34. Kf2 Bf5
	// 35. Ra7 g6 36. Ra6+ Kc5 37. Ke1 Nf4 38. g3 Nxh3 39. Kd2 Kb5 40. Rd6 Kc5 41. Ra6
	// Nf2 42. g4 Bd3 43. Re6 1/2-1/2

	var moves []string

	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()
		log.Printf(" token=%q, inComment=%v, strings.HasSuffix=%v", token, state.inComment, strings.HasSuffix(token, "}"))

		if state.inComment {
			state.inComment = !closingComment(token)
			continue
		}

		if openingComment(token) {
			if state.inComment {
				return nil, fmt.Errorf("comments do not nest!")
			}
			state.inComment = true
			continue
		}

		if strings.HasSuffix(token, "...") {
			// like: 2. Ba4 {This is a comment.} 2... Nf6
			continue
		}

		if strings.HasSuffix(token, ".") {
			// sanity check
			num := token[0 : len(token)-1]
			gotMoveIdx, err := strconv.Atoi(num)
			if err != nil || state.moveIdx != gotMoveIdx {
				return nil, fmt.Errorf("expected moveIdx %d, got %d from %q", state.moveIdx, gotMoveIdx, num)
			}
		} else {
			moves = append(moves, token)
			if state.turn == Black {
				state.moveIdx++
			}
			state.turn = state.turn.Opponent()
		}

	}
	return moves, nil
}

func openingComment(token string) bool {
	return strings.HasPrefix(token, "{") || strings.HasPrefix(token, "(")
}

func closingComment(token string) bool {
	return strings.HasSuffix(token, "}") || strings.HasSuffix(token, ")")
}
