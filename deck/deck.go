package deck

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Card struct {
	Suit  string
	Rank  string
	Value int
}

var r *rand.Rand

const maxJokers = 13

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func New(sortMethod string, numberOfDecks int, filters []string) []Card {
	if numberOfDecks < 1 {
		numberOfDecks = 1
	}

	var deck []Card

	for a := 0; a < numberOfDecks; a++ {
		for j := 0; j < 4; j++ {
			var suit string
			switch j {
			case 0:
				// suit = "spade"
				suit = "\u2660"
				break
			case 1:
				suit = "\u2666"
				break
			case 2:
				suit = "\u2663"
				break
			case 3:
				suit = "\u2665"
				break
			}

			for i := 1; i < 14; i++ {
				var rank string
				switch i {
				case 1:
					rank = "A"
					break
				case 11:
					rank = "J"
					break
				case 12:
					rank = "Q"
					break
				case 13:
					rank = "K"
					break
				default:
					rank = strconv.FormatInt(int64(i), 10)
				}

				card := Card{
					Suit:  suit,
					Rank:  rank,
					Value: i,
				}
				deck = append(deck, card)
			}
		}
	}

	switch sortMethod {
	case "byRank":
		sort.SliceStable(deck, func(i, j int) bool {
			return deck[i].Value < deck[j].Value
		})
		break
	case "bySuit":
		sort.SliceStable(deck, func(i, j int) bool {
			return deck[i].Suit < deck[j].Suit
		})
		break
	case "shuffle":
		deck = shuffle(deck)
		break
	case "addJockers":
		jokers := r.Intn(maxJokers)
		for i := 0; i < jokers; i++ {
			joker := Card{Suit: "", Rank: "joker"}
			deck = append(deck, joker)
		}
		break
	}

	for _, filter := range filters {
		var tempDeck []Card
		for _, card := range deck {
			if card.Rank == filter {
				continue
			}
			tempDeck = append(tempDeck, card)
		}
		deck = tempDeck
	}

	return deck
}

func shuffle(deck []Card) []Card {
	cards := len(deck)
	tempDeck := make([]Card, cards)

	myRand := r.Perm(cards)
	for i := range myRand {
		tempDeck[i] = deck[i]
	}

	return tempDeck
}
