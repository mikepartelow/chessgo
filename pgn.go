package chessgo

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type PGNTag struct {
	Tag   string
	Value string
}

func ParsePGN(in io.Reader) ([]PGNTag, error) {
	var tags []PGNTag

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()
		if tag := parseTag(line); tag != nil {
			tags = append(tags, *tag)
		}
	}

	if len(tags) > 0 {
		return tags, nil
	}

	return nil, fmt.Errorf("invalid input")
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
