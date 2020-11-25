package go_pgn

import (
	"fmt"
	"strings"
)

const tagFmtString = "[%s \"%s\"]\n"

func exportFormat(p *PGNGame) string {
	var b strings.Builder
	writeHeaders(&b, p)
	writeMoves(&b, p.Moves, 1)
	fmt.Fprintf(&b, " %s", p.Result.String())
	fmt.Fprintln(&b)
	return b.String()
}

func writeHeaders(b *strings.Builder, p *PGNGame) {
	if p.Event != "" {
		fmt.Fprintf(b, tagFmtString, "Event", p.Event)
	}
	if p.Site != "" {
		fmt.Fprintf(b, tagFmtString, "Site", p.Site)
	}
	if p.Date != "" {
		fmt.Fprintf(b, tagFmtString, "Date", p.Date)
	}
	if p.Round != "" {
		fmt.Fprintf(b, tagFmtString, "Round", p.Round)
	}
	if p.White != "" {
		fmt.Fprintf(b, tagFmtString, "White", p.White)
	}
	if p.Black != "" {
		fmt.Fprintf(b, tagFmtString, "Black", p.Black)
	}

	if p.Result != GameResultNotPresent {
		fmt.Fprintf(b, tagFmtString, "Result", p.Result.String())
	}

	for _, e := range p.OptionalTags {
		fmt.Fprintf(b, tagFmtString, e.k, e.v)
	}

	fmt.Fprintf(b, "\n")
}

func writeMoves(b *strings.Builder, m *Move, moveNumber int) {
	if m == nil {
		return
	}

	if moveNumber > 1 {
		fmt.Fprint(b, " ")
	}

	fmt.Fprintf(b, "%d.", moveNumber)
	fmt.Fprintf(b, " %s", m.WhiteSAN)

	blackMoveNumberIndication := false

	if m.WhiteComment != "" {
		fmt.Fprintf(b, " {%s}", m.WhiteComment)
		blackMoveNumberIndication = true
	}
	if nag, ok := m.WhiteNAG.Get(); ok {
		fmt.Fprintf(b, " $%d", nag)
	}

	vardone := make([]bool, len(m.Variation))

	// if the variation was started by white, then we put the variation in
	// between the move and have black's play be 4... Kc3
	for i, v := range m.Variation {
		if v.WhiteSAN != "" {
			fmt.Fprint(b, " (")
			writeMoves(b, v, moveNumber)
			fmt.Fprint(b, ")")
			blackMoveNumberIndication = true
			vardone[i] = true
		}
	}

	if m.BlackSAN == "" {
		return
	}

	if blackMoveNumberIndication {
		fmt.Fprintf(b, " %d...", moveNumber)
	}

	fmt.Fprintf(b, " %s", m.BlackSAN)
	if m.BlackComment != "" {
		fmt.Fprintf(b, " {%s}", m.BlackComment)
	}
	if nag, ok := m.BlackNAG.Get(); ok {
		fmt.Fprintf(b, " $%d", nag)
	}

	// if started by black, we put the variation at the end of the move chain
	// start with (4... Kc3)
	for i, v := range m.Variation {
		if !vardone[i] {
			fmt.Fprintf(b, " (%d...", moveNumber)
			fmt.Fprintf(b, " %s", v.BlackSAN)
			if v.BlackComment != "" {
				fmt.Fprintf(b, " {%s}", v.BlackComment)
			}
			if nag, ok := v.BlackNAG.Get(); ok {
				fmt.Fprintf(b, " $%d", nag)
			}
			writeMoves(b, v.Next, moveNumber+1)
			fmt.Fprint(b, ")")
		}
	}

	writeMoves(b, m.Next, moveNumber+1)
}
