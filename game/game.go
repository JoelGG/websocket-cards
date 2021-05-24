package game

type GameController struct {
	moves    chan Incoming
	response chan Outcoming
	state    GameState
	history  []GameState
}

type GameState struct {
	players map[int]Player
	index   int
	deck    []Card
}

type Player struct {
	hand []Card
}

type Move struct {
	Player   int
	MoveType int
}

type Incoming struct {
	Move      Move
	GameState GameState
}

type Outcoming struct {
	Success   bool
	GameState GameState
}

func NewGameController(playercount int, incoming chan Incoming, response chan Outcoming) *GameController {
	g := GameController{}
	g.moves = incoming
	g.response = response
	g.state = NewGameState()
	return &g
}

func NewGameState() GameState {
	g := GameState{}
	g.players = map[int]Player{}
	g.deck = NewDeck()
	return g
}

func NewPlayer() Player {
	p := Player{}
	p.hand = []Card{}
	return p
}

func (g *GameController) Start() {
	for {
		inc := <-g.moves
		if inc.Move.Player != g.state.index {
			g.response <- Outcoming{Success: false}
		} else {
			g.history = append(g.history, g.state)
			g.state = g.state.Consume(inc.Move)
			g.response <- Outcoming{Success: true, GameState: g.state}
		}
	}
}

func (g GameState) Consume(m Move) GameState {
	g.index++
	return g
}
