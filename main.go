package go_pgn

type PGNDatabase struct {
	games []PGNGame
}

type TagKVPair struct {
	k string
	v string
}

type PGNGame struct {
	Event        string
	Site         string
	Date         string
	Round        string
	White        string
	Black        string
	Result       GameResult
	OptionalTags []TagKVPair
	Moves        *Move
}

type Move struct {
	WhiteSAN     string
	WhiteComment string
	WhiteNAG     OptionalUint8
	BlackSAN     string
	BlackComment string
	BlackNAG     OptionalUint8
	Next         *Move
	Variation    []*Move
}

type GameResult int

const (
	GameResultNotPresent GameResult = iota
	GameResultDraw
	GameResultWhite
	GameResultBlack
	GameResultOther
)

func (g GameResult) String() string {
	switch g {
	case GameResultDraw:
		return "1/2-1/2"
	case GameResultWhite:
		return "1-0"
	case GameResultBlack:
		return "0-1"
	case GameResultNotPresent:
		return ""
	case GameResultOther:
		fallthrough
	default:
		return "*"
	}
}

type OptionalUint8 struct {
	set bool
	val uint8
}

func (o *OptionalUint8) Set(i uint8) {
	o.set = true
	o.val = i
}

func (o *OptionalUint8) Clear() {
	o.set = false
}

func (o *OptionalUint8) Get() (uint8, bool) {
	return o.val, o.set
}

func (p *PGNGame) String() string {
	return exportFormat(p)
}

func Validate(pgnString string) (*PGNGame, error) {
	return nil, nil
}
