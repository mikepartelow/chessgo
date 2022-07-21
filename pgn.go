package chessgo

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type PGNTag struct {
	Tag   string
	Value string
}

func ParsePGN(in io.Reader) ([]PGNTag, []string, error) {
	p := NewPGNParser(in)

	err := p.Parse()
	if err != nil {
		return nil, nil, fmt.Errorf("problem parsing PGN: %v", err)
	}

	return p.Tags(), p.Moves(), nil
}

type PGNParser struct {
	scanner   *bufio.Scanner
	inComment bool
	turn      Color
	moveIdx   int

	tags  []PGNTag
	moves []string
}

func NewPGNParser(in io.Reader) *PGNParser {
	return &PGNParser{
		scanner: bufio.NewScanner(in),
		moveIdx: 1,
	}
}

func (p *PGNParser) Parse() error {
	for p.scanner.Scan() {
		line := p.scanner.Text()
		if isBracketedWith("[", "]", line) {
			if err := p.parseTag(line); err != nil {
				return fmt.Errorf("problem parsing tags: %v", err)
			}
		} else {
			if err := p.parseMoves(line); err != nil {
				return fmt.Errorf("problem parsing moves: %v", err)
			}
		}
	}

	return nil
}

func (p *PGNParser) Tags() []PGNTag {
	return p.tags
}

func (p *PGNParser) Moves() []string {
	return p.moves
}

func (p *PGNParser) parseTag(line string) error {
	parts := strings.SplitN(line[1:len(line)-1], " ", 2)

	if len(parts) == 2 && isBracketedWith(`"`, `"`, parts[1]) {
		tag := parts[0]
		value := parts[1][1 : len(parts[1])-1]
		p.tags = append(p.tags, PGNTag{tag, value})
		return nil
	}

	return fmt.Errorf("could not parse (alleged) tag: %s", line)
}

func (p *PGNParser) parseMoves(line string) error {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()
		// log.Printf(" token=%q, inComment=%v, strings.HasSuffix=%v", token, p.inComment, strings.HasSuffix(token, "}"))

		if p.inComment {
			p.inComment = !closingComment(token)
		} else if openingComment(token) {
			if p.inComment {
				return fmt.Errorf("PGN comments do not nest!")
			}
			p.inComment = true
		} else if strings.HasSuffix(token, "...") {
			// like: 2. Ba4 {This is a comment.} 2... Nf6
			// consume and do nothing
		} else if strings.HasSuffix(token, ".") {
			// sanity check
			num := token[0 : len(token)-1]
			gotMoveIdx, err := strconv.Atoi(num)
			if err != nil || p.moveIdx != gotMoveIdx {
				return fmt.Errorf("expected moveIdx %d, got %d from %q", p.moveIdx, gotMoveIdx, num)
			}
		} else {
			p.moves = append(p.moves, token)
			if p.turn == Black {
				p.moveIdx++
			}
			p.turn = p.turn.Opponent()
		}
	}

	return nil
}

func isBracketedWith(leftBracket, rightBracket, text string) bool {
	return strings.HasPrefix(text, leftBracket) && strings.HasSuffix(text, rightBracket)
}

func openingComment(token string) bool {
	return strings.HasPrefix(token, "{") || strings.HasPrefix(token, "(")
}

func closingComment(token string) bool {
	return strings.HasSuffix(token, "}") || strings.HasSuffix(token, ")")
}
