package game

type Game struct {
	player Player
	dealer Dealer
	deck   Deck
}

type Cards []Card

type Card struct {
	group string
	face  string
	value int
}

type Player struct {
	cards Cards
}

type Dealer struct {
	cards Cards
}

type Deck struct {
	cards [52]Card
}

func NewDeck() Deck {
	// groups := [4]string{"S", "D", "C", "H"}
	// faces := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "K", "Q", "A", "J"}
	
	// all_cards := [52]Card{}
	return *new(Deck)

}

func NewGame() Game {
	cards := Cards{Card{"S","2",2}, Card{"S","2",2}}
	player := Player{cards}
	dealer := Dealer{cards}
	game := Game{player, dealer, *new(Deck)}
	return game
}

// func (deck Deck) Draw() Card {
// 	return
// }

func (game Game) NumberOfCardsForPlayer() int {
	return len(game.player.cards)
}

func (game Game) NumberOfCardsForDealer() int {
	return len(game.dealer.cards)
}

// func (game Game) start() int {
// 	game.player.addCard(game.deck.draw())
// 	game.player.addCard(game.deck.draw())
// 	game.dealer.addCard(game.deck.draw())
// 	game.dealer.addCard(game.deck.draw())
// }
