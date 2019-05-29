package mrctech

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Suit int
	Number int
}

type Deck struct {
	Cards []Card
}

func GenerateStraightFlush() Deck{
	rand.Seed(time.Now().UnixNano())
	sameSuit := rand.Intn(4)
	straight := rand.Intn(9)+1

	c1 := Card{sameSuit, straight}
	c2 := Card{sameSuit, straight+1}
	c3 := Card{sameSuit, straight+2}
	c4 := Card{sameSuit, straight+3}
	c5 := Card{sameSuit, straight+4}

	if straight == 10{
		c5 = Card{sameSuit, 1}
	}

	var deck []Card
	deck = append(deck, c1, c2, c3, c4, c5)
	return Deck{deck}
}

func GenerateFourOfAKind() Deck{
	rand.Seed(time.Now().UnixNano())
	suit := rand.Intn(4)
	numb := rand.Intn(13)+1

	c1 := Card{suit, numb}
	c2 := Card{suit, numb}
	c3 := Card{suit, numb}
	c4 := Card{suit, numb}
	c5 := Card{suit, numb}


}

func GenerateFullHouse() Deck{
	rand.Seed(time.Now().UnixNano())
}

func GenerateFlush() Deck{
	rand.Seed(time.Now().UnixNano())
}

func GenerateStraight() Deck{
	rand.Seed(time.Now().UnixNano())
}

func GenerateUnmatch() Deck{
	rand.Seed(time.Now().UnixNano())
}

func GenerateCard() Card{

}

func GenerateDeck() Deck{

}

// Return suit numbers. For example, ♣♡♠♢
func (d Deck)SuitLevel() (int, int){
	suits := make(map[int]int)
	for _, card := range d.Cards{
		c, ok := suits[card.Suit]
		if !ok {
			suits[card.Suit] = 0
		}else{
			suits[card.Suit] = c+1
		}
	}

	maxSuit := 0
	max := 0
	for i, c := range suits{
		if c > max {
			maxSuit = i
			max = c
		}
	}

	return maxSuit, max
}

func (d Deck)IsStraight() bool{
	var numbs []int
	for _, card := range d.Cards{
		numbs = append(numbs, card.Number)
	}
	sort.Ints(numbs)

	sub := numbs[4]-numbs[0]
	// Check regular straight if sub == 4, Check A, T, J, Q, K if sub ==12 && sub2 == 3
	if sub == 4{
		return true
	}else if sub == 12{
		if numbs[4]-numbs[1] == 3{
			return true
		}
	}

	return false
}

func (d Deck)DisplayDeck() {
	displayListOfSuit := []string{"♣", "♢", "♡", "♠"}
	displayListOfNumb := []string{"X", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	fmt.Print("Deck: ")
	for _, card := range d.Cards {
		fmt.Print(displayListOfSuit[card.Suit], displayListOfNumb[card.Number], ", ")
	}
	fmt.Print("\n")
}

func (d Deck)CompareTo(d2 Deck) int{

	return 0
}