package game

type Card int

func (c Card) GetSuit() int {
	return int(c) % 13
}

func (c Card) GetIndex() int {
	return int(c) % 13
}

func NewDeck() []Card {
	deck := []Card{}
	for i := 0; i < 52; i++ {
		deck = append(deck, Card(i))
	}
	return deck
}
