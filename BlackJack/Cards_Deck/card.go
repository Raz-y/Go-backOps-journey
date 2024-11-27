//go:generate stringer -type=Suit,Rank

/*
Package cardsdeck provides utilities for creating, managing, and manipulating decks of playing cards.
It supports standard deck operations such as shuffling, sorting, adding Jokers, and filtering cards.
*/
package cardsdeck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents the suit of a card (Spade, Diamond, Club, Heart, Joker).
type Suit uint8

const (
	Spade Suit = iota // 0
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank represents the rank of a card (Ace, Two, ..., King).
type Rank uint8

const (
	_ Rank = iota // 0
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank  = Ace
	maxRank  = King
	numRanks = 13
)

// Card represents a playing card with a Suit, Rank, and an internal absolute rank for sorting.
type Card struct {
	Suit
	Rank
	absRank int // Cached absolute rank for optimized sorting
}

// String returns a human-readable string representation of a Card.
// Example: "Ace of Spades" or "Joker".
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates a new deck of cards with the specified options.
// Options can modify the deck, such as adding Jokers or shuffling.
//
// Example:
//
// cards := New(Jokers(2), Shuffle)
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{
				Suit:    suit,
				Rank:    rank,
				absRank: int(suit)*numRanks + int(rank),
			})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards

}

// DefaultSort sorts a deck of cards in the default order (Spades, Diamonds, Clubs, Hearts, sorted by rank).
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort returns a sorting function that uses the provided comparison function for custom sorting.
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	// Returns a comparison function for sorting cards by their absolute rank.
	return func(i, j int) bool {
		return cards[i].absRank < cards[j].absRank
	}
}

// Shuffle randomizes the order of the cards in the deck.
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers adds the specified number of Jokers to the deck.
// Panics if n is negative.
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		if n < 0 {
			panic("Jokers count cannot be negative")
		}
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank:    Rank(i),
				Suit:    Joker,
				absRank: -1, // Special rank for Jokers
			})
		}
		return cards
	}
}

// Filter removes cards from the deck that match the specified condition.
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(card []Card) []Card {
		var ret []Card
		for _, c := range card {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// Deck duplicates the deck n times.
// Panics if n is less than 1.
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		if n < 1 {
			panic("Deck count must be at least 1")
		}
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

// Equals checks if two cards are identical by comparing their Suit and Rank.
func (c Card) Equals(other Card) bool {
	return c.Suit == other.Suit && c.Rank == other.Rank
}
