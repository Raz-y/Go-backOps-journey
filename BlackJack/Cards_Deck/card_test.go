package cardsdeck

import (
	"fmt"
	"testing"
)

// ExampleCard demonstrates the string representation of cards.
// It shows how to print a variety of cards including a Joker.
func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Nine, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	//Output:
	//Ace of Hearts
	//Two of Clubs
	//Nine of Spades
	//Joker
}

// TestNew ensures that a new deck contains the correct number of cards.
func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Error("Wrong number of cards")
	}
}

// TestDefaultSort verifies that DefaultSort arranges the deck
// in ascending order by suit and rank.
func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade, absRank: int(Spade)*numRanks + int(Ace)}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

// TestSort checks custom sorting using the Less function.
func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade, absRank: int(Spade)*numRanks + int(Ace)}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

// TestJokers ensures that the specified number of Jokers is added to the deck.
func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0

	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, received: ", count)
	}
}

// TestFilter ensures that cards matching the provided filter are removed from the deck.
func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

// TestDeck verifies that the Deck function correctly duplicates the deck n times.
func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, received %d cards.", 13*4*3, len(cards))
	}
}

// TestEquals checks that the Equals method correctly identifies equal and unequal cards.
func TestEquals(t *testing.T) {
	card1 := Card{Rank: Ace, Suit: Heart}
	card2 := Card{Rank: Ace, Suit: Heart}
	card3 := Card{Rank: Two, Suit: Heart}

	if !card1.Equals(card2) {
		t.Error("Expected cards to be equal")
	}
	if card1.Equals(card3) {
		t.Error("Expected cards to be unequal")
	}
}

// TestInvalidJokers ensures that Jokers panics when given a negative count.
func TestInvalidJokers(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for negative jokers")
		}
	}()
	New(Jokers(-1))
}

// TestInvalidDeck ensures that Deck panics when given a count less than 1.
func TestInvalidDeck(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero decks")
		}
	}()
	New(Deck(0))
}

// TestShuffleRandomness verifies that Shuffle produces a deck in a different order
// from the original deck.
func TestShuffleRandomness(t *testing.T) {
	cards1 := New()
	cards2 := New(Shuffle)

	identical := true
	for i, card := range cards1 {
		if !card.Equals(cards2[i]) {
			identical = false
			break
		}
	}

	if identical {
		t.Error("Expected shuffled deck to differ from original")
	}
}
