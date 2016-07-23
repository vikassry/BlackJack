package game

import (
	"testing"
)

func TestBlackJackGameShouldStart(t *testing.T) {
	cards := Cards{Card{"s", "1"}, Card{"s", "2"}}
	player := Player{cards}
	dealer := Dealer{cards}
	deck := Deck{}
	game := Game{player, dealer, deck}
	players_card := game.NumberOfCardsForPlayer()
	dealers_card := game.NumberOfCardsForDealer()

	if players_card != 2 {
		t.Errorf("Expected number of cards to be 2 for player, got %d", players_card)
	}
	if dealers_card != 2 {
		t.Errorf("Expected number of cards to be 2 for dealer, got %d", dealers_card)
	}
}
