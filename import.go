package go_pgn

import "unicode"

func importPGN(s string) *PGNGame {
	return nil
}

func importHeaders(s string, p *PGNGame) {

}

func importGame(s string, p *PGNGame) {

}

func tokenize(s string) {
}

func skipSpaces(i int, s string) int {
	// skip spaces
	for ; i < len(s); i++ {
		if !unicode.IsSpace(rune(s[i])) {
			break
		}
	}
	return i
}

func expectTagKey(i int, s string) (string, int) {
}

func expectTagValue(i int, s string) (string, int) {
}

func tokenizeHeader(i int, s string) ([]string, int) {
	tokens := []string{}

	var key, val string
	for ; i < len(s); i++ {
		i = skipSpaces(i, s)

		if s[i] != '['

		key, i = expectTagKey(i, s)

		i = skipSpaces(i, s)
		value, i = expectTagValue(i, s)

		if s[i] == '[' || s[i] == ']' {
			tokens = append(tokens, string(s[i]))
		}
	}

	return tokens, i
}
