package go_pgn

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	moves := &Move{
		WhiteSAN: "e4",
		BlackSAN: "e5",
		Next: &Move{
			WhiteSAN: "Nf3",
			BlackSAN: "Nf6",
			Variation: []*Move{
				&Move{
					BlackSAN: "d5",
					Next: &Move{
						WhiteSAN: "a3",
						BlackSAN: "h6",
					},
				},
			},
			BlackComment: "Petrov's defense",
			Next: &Move{
				WhiteSAN: "a3",
				BlackSAN: "h6",
			},
		},
	}

	g := &PGNGame{
		Event:  "test event",
		Site:   "USA",
		Date:   "2020.3.11",
		Round:  "7",
		White:  "white player",
		Black:  "black player",
		Result: GameResultDraw,
		Moves:  moves,
	}

	expect := `[Event "test event"]
[Site "USA"]
[Date "2020.3.11"]
[Round "7"]
[White "white player"]
[Black "black player"]
[Result "1/2-1/2"]

1. e4 e5 2. Nf3 Nf6 {Petrov's defense} (2... d5 3. a3 h6) 3. a3 h6 1/2-1/2
`

	assert.Equal(t, expect, g.String())
}

func TestExportHeadersEmpty(t *testing.T) {
	var b strings.Builder
	expect := "\n"
	g := &PGNGame{}
	writeHeaders(&b, g)
	assert.Equal(t, expect, b.String())
}

func TestExportHeadersAll(t *testing.T) {
	var b strings.Builder
	expect := `[Event "test event"]
[Site "USA"]
[Date "2020.3.11"]
[Round "7"]
[White "white player"]
[Black "black player"]
[Result "1/2-1/2"]

`
	g := &PGNGame{
		Event:  "test event",
		Site:   "USA",
		Date:   "2020.3.11",
		Round:  "7",
		White:  "white player",
		Black:  "black player",
		Result: GameResultDraw,
	}

	writeHeaders(&b, g)
	assert.Equal(t, expect, b.String())
}

func TestExportHeadersWithOptional(t *testing.T) {
	var b strings.Builder
	expect := `[Event "test event"]
[Site "USA"]
[Date "2020.3.??"]
[Round "?"]
[White "????"]
[Black "????"]
[Result "1/2-1/2"]
[WhiteElo "1450"]
[BlackElo "1600"]
[Website "chess.com"]

`
	g := &PGNGame{
		Event:  "test event",
		Site:   "USA",
		Date:   "2020.3.??",
		Round:  "?",
		White:  "????",
		Black:  "????",
		Result: GameResultDraw,
		OptionalTags: []TagKVPair{
			{k: "WhiteElo", v: "1450"},
			{k: "BlackElo", v: "1600"},
			{k: "Website", v: "chess.com"},
		},
	}

	writeHeaders(&b, g)
	assert.Equal(t, expect, b.String())
}
