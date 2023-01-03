package deck

import (
	"fmt"
	"math/rand"
	"time"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Three, Suit: Diamond})
	fmt.Println(Card{Rank: Seven, Suit: Club})
	fmt.Println(Card{Rank: Queen, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Three of Diamonds
	// Seven of Clubs
	// Queen of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if (cards[0] != expected) {
		t.Error("Expected Ace of Spades as first card. Recieved:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades as first card. Recieved:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	numJokers := r.Intn(4)
	cards := New(Jokers(numJokers))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != numJokers {
		t.Error("Expected", numJokers, "Jokers, recieved:", count)
	}
}

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

func TestDeck(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	numDecks := r.Intn(5) + 1
	cards := New(Deck(numDecks))
	if len(cards) != 52 * numDecks {
		t.Error("Expected", 52*numDecks, "cards, recieved:", len(cards))
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	original := New()
	first := original[40]
	second:= original[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, recieved: %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be %s, recieved: %s", second, cards[1])
	}
}