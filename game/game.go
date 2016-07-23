package game

type Game struct {
	player Player
	dealer Dealer
	deck   Deck
}

type Cards []Card

type Card struct {
	value string
	group string
}

type Player struct {
	cards Cards
}

type Dealer struct {
	cards Cards
}

type Deck struct {
	cards Cards
}

func (game Game) NumberOfCardsForPlayer() int {
	return len(game.player.cards)
}

func (game Game) NumberOfCardsForDealer() int {
	return len(game.dealer.cards)
}
