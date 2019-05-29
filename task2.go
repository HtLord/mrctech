package mrctech

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Suit   int
	Number int
}

type Deck struct {
	Cards []Card
}

func GenerateStraightFlush() Deck {
	rand.Seed(time.Now().UnixNano())
	sameSuit := rand.Intn(4)
	straight := rand.Intn(9) + 1

	c1 := Card{sameSuit, straight}
	c2 := Card{sameSuit, straight + 1}
	c3 := Card{sameSuit, straight + 2}
	c4 := Card{sameSuit, straight + 3}
	c5 := Card{sameSuit, straight + 4}

	if straight == 10 {
		c5 = Card{sameSuit, 1}
	}

	var deck []Card
	deck = append(deck, c1, c2, c3, c4, c5)
	return Deck{deck}
}

func GenerateFourOfAKind() Deck {
	rand.Seed(time.Now().UnixNano())
	suit := rand.Intn(4)
	numb := rand.Intn(13) + 1
	numb2 := 0

	for {
		numb2 = rand.Intn(13) + 1
		if numb2 != numb {
			break
		}
	}

	c1 := Card{0, numb}
	c2 := Card{1, numb}
	c3 := Card{2, numb}
	c4 := Card{3, numb}
	c5 := Card{suit, numb2}

	var deck []Card
	deck = append(deck, c1, c2, c3, c4, c5)
	return Deck{deck}
}

func GenerateFullHouse() Deck {
	rand.Seed(time.Now().UnixNano())
	suit := rand.Intn(4)
	suit2 := rand.Intn(4)
	numb := rand.Intn(13) + 1
	numb2 := 0

	for {
		numb2 = rand.Intn(13) + 1
		if numb2 != numb {
			break
		}
	}

	c1 := Card{suit % 4, numb}
	c2 := Card{(suit + 1) % 4, numb}
	c3 := Card{(suit + 2) % 4, numb}

	c4 := Card{suit2 % 4, numb2}
	c5 := Card{(suit2 + 1) % 4, numb2}

	var deck []Card
	deck = append(deck, c1, c2, c3, c4, c5)
	return Deck{deck}
}

func GenerateFlush() Deck {
	rand.Seed(time.Now().UnixNano())
	suit := rand.Intn(4)
	cards := []Card{}
	seqTimes := 0
	seq := 0
	for {
		if len(cards) == 5 {
			break
		}

		if seqTimes == 4 {
			seq += 2
			cards = append(cards, Card{suit, seq%13 + 1})
			break
		}

		if rand.Intn(2) == 0 {
			seqTimes++
			seq++
			cards = append(cards, Card{suit, seq%13 + 1})
		} else {
			seq += 2
			cards = append(cards, Card{suit, seq%13 + 1})
		}
	}

	return Deck{cards}
}

func GenerateStraight() Deck {
	rand.Seed(time.Now().UnixNano())
	straight := rand.Intn(9) + 1
	suits := map[int]int{0: 0, 1: 0, 2: 0, 3: 0}

	suit := rand.Intn(4)
	suits[suit] = suits[suit] + 1
	c1 := Card{suit, straight}
	suit = rand.Intn(4)
	suits[suit] = suits[suit] + 1
	c2 := Card{suit, straight + 1}
	suit = rand.Intn(4)
	suits[suit] = suits[suit] + 1
	c3 := Card{suit, straight + 2}
	suit = rand.Intn(4)
	suits[suit] = suits[suit] + 1
	c4 := Card{suit, straight + 3}
	suit = rand.Intn(4)
	if suits[suit] == 4 {
		suit = (suit + 1) % 4
	}
	c5 := Card{suit, straight + 4}

	if straight == 10 {
		c5 = Card{suit, 1}
	}

	var deck []Card
	deck = append(deck, c1, c2, c3, c4, c5)
	return Deck{deck}
}

// Return suit numbers. For example, ♣♡♠♢
func (d Deck) SuitLevel() (int, int) {
	suits := make(map[int]int)
	for _, card := range d.Cards {
		c, ok := suits[card.Suit]
		if !ok {
			suits[card.Suit] = 0
		} else {
			suits[card.Suit] = c + 1
		}
	}

	maxSuit := 0
	max := 0
	for i, c := range suits {
		if c > max {
			maxSuit = i
			max = c
		}
	}

	return maxSuit, max
}

func SortDeckToNumber(d Deck) []int {
	a := []int{}
	for _, c := range d.Cards {
		a = append(a, c.Number)
	}
	sort.Ints(a)
	return a
}

func (d Deck) MaxNumberMatch() int {
	numb := 0
	max := 1
	c := 1
	a := SortDeckToNumber(d)

	numb = a[0]
	for i := 1; i < len(a); i++ {
		if numb == a[i] {
			c++
			if c > max {
				max = c
			}
		} else {
			c = 1
			numb = a[i]
		}
	}

	return max
}

func (d Deck) IsFourOfAKind() bool {
	return d.MaxNumberMatch() == 4
}

func (d Deck) IsTri() bool {
	return d.MaxNumberMatch() == 3
}

func (d Deck) IsFullHouse() bool {
	a := SortDeckToNumber(d)
	n := -1
	change := 0
	for _, v := range a {
		if n < 0 {
			n = v
		}

		if n != v {
			change++
			n = v
		}
	}

	return change == 2 && !d.IsFourOfAKind()
}

func (d Deck) IsStraightFlush() bool {
	return d.IsStraight() && d.IsFlush()
}

func (d Deck) IsFlush() bool {
	suit := -1
	for _, card := range d.Cards {
		if suit >= 0 {
			if suit != card.Suit {
				return false
			}
		} else {
			suit = card.Suit
		}
	}
	return true
}

func (d Deck) IsStraight() bool {
	var numbs []int
	for _, card := range d.Cards {
		numbs = append(numbs, card.Number)
	}
	sort.Ints(numbs)

	sub := numbs[4] - numbs[0]
	// Check regular straight if sub == 4, Check A, T, J, Q, K if sub ==12 && sub2 == 3
	if sub == 4 {
		return true
	} else if sub == 12 {
		if numbs[4]-numbs[1] == 3 {
			return true
		}
	}

	return false
}

func (d Deck) DisplayDeck() {
	displayListOfSuit := []string{"♣", "♢", "♡", "♠"}
	displayListOfNumb := []string{"X", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	fmt.Print("Deck: ")
	for _, card := range d.Cards {
		fmt.Print(displayListOfSuit[card.Suit], displayListOfNumb[card.Number], ", ")
	}
	fmt.Print("\n")
}

// return 0 = equal; 1 = d1 bigger; 2 = d1 smaller
func (d Deck) CompareTo(d2 Deck) int {
	return 0
}
