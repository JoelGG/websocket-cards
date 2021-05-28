package game

import (
	"encoding/json"
	"fmt"
)

type GameController interface {
	Consume(Incoming) (Outgoing, error)
}

type Incoming struct {
	Player int
	Msg    string
}

type Outgoing struct {
	Msg string
}

type GameControllerBrag struct {
	history []GameStateBrag
	state   GameStateBrag
}

type GameControllerEcho struct{}

func (g *GameControllerEcho) Consume(inc Incoming) (Outgoing, error) {
	return Outgoing{Msg: inc.Msg}, nil
}

func NewGameControllerBrag(playercount int) *GameControllerBrag {
	g := GameControllerBrag{}
	g.history = []GameStateBrag{}
	g.state = NewGameState()
	return &g
}

func (g *GameControllerBrag) Consume(inc Incoming) (Outgoing, error) {
	move, err := g.msgDecode(inc.Msg)

	if err != nil {
		return Outgoing{}, err
	}

	newstate, err := g.state.Consume(move)
	fmt.Println(newstate.Index)

	if err != nil {
		return Outgoing{}, err
	} else {
		g.history = append(g.history, g.state)
		g.state = newstate
		otg, err := g.stateEncode(g.state)

		if err != nil {
			return Outgoing{}, err
		}

		return Outgoing{otg}, nil
	}
}

func (g *GameControllerBrag) stateEncode(state GameStateBrag) (string, error) {
	otp, err := json.Marshal(state)

	fmt.Println(state.Deck)

	if err != nil {
		return "", err
	} else {
		return string(otp), nil
	}
}

func (g *GameControllerBrag) msgDecode(msg string) (Move, error) {
	mv := &Move{}
	err := json.Unmarshal([]byte(msg), mv)

	if err != nil {
		return Move{}, err
	} else {
		return *mv, nil
	}
}
