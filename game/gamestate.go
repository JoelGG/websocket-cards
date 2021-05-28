package game

type GameStateBrag struct {
	Players []Player `json:"players"`
	Index   int      `json:"index"`
	Deck    []Card   `json:"deck"`
}

type Player struct {
	Hand []Card `json:"hand"`
}

type Move struct {
	Player   int `json:"player"`
	MoveType int `json:"movetype"`
}

func NewGameState() GameStateBrag {
	g := GameStateBrag{}
	g.Players = []Player{}
	g.Deck = NewDeck()
	g.Index = 0
	return g
}

func NewPlayer() Player {
	p := Player{}
	p.Hand = []Card{}
	return p
}

func (g *GameStateBrag) Consume(m Move) (GameStateBrag, error) {
	g.Index++
	return *g, nil
}
