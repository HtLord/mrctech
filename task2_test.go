package mrctech

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestGenerateStraightFlush(t *testing.T) {
//	for i := 0; i < 5; i++ {
//		GenerateStraightFlush().DisplayDeck()
//	}
//}
//
//func TestGenerateFourOfAKind(t *testing.T) {
//	for i := 0; i < 5; i++ {
//		GenerateFourOfAKind().DisplayDeck()
//	}
//}
//
//func TestGenerateFullHouse(t *testing.T) {
//	for i := 0; i < 5; i++ {
//		GenerateFixFullHouse().DisplayDeck()
//	}
//}
//
//func TestGenerateFlush(t *testing.T) {
//	for i := 0; i < 5; i++ {
//		GenerateFlush().DisplayDeck()
//	}
//}
//
//func TestGenerateStraight(t *testing.T) {
//	for i := 0; i < 5; i++ {
//		GenerateFixSuitStraight().DisplayDeck()
//	}
//}

func TestDeck_IsFlush(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()

	assert.Equal(t, true, sf.IsFlush())
	assert.Equal(t, false, fk.IsFlush())
	assert.Equal(t, false, fh.IsFlush())
	assert.Equal(t, true, f.IsFlush())
	assert.Equal(t, false, s.IsFlush())
}

func TestDeck_IsStraight(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()

	assert.Equal(t, true, sf.IsStraight())
	assert.Equal(t, false, fk.IsStraight())
	assert.Equal(t, false, fh.IsStraight())
	assert.Equal(t, false, f.IsStraight())
	assert.Equal(t, true, s.IsStraight())
}

func TestDeck_IsStraightFlush(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()

	assert.Equal(t, true, sf.IsStraightFlush())
	assert.Equal(t, false, fk.IsStraightFlush())
	assert.Equal(t, false, fh.IsStraightFlush())
	assert.Equal(t, false, f.IsStraightFlush())
	assert.Equal(t, false, s.IsStraightFlush())
}

func TestDeck_IsFourOfAKind(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()

	assert.Equal(t, false, sf.IsFourOfAKind())
	assert.Equal(t, true, fk.IsFourOfAKind())
	assert.Equal(t, false, fh.IsFourOfAKind())
	assert.Equal(t, false, f.IsFourOfAKind())
	assert.Equal(t, false, s.IsFourOfAKind())
}

func TestDeck_IsFullHouse(t *testing.T) {
	sf := GenerateStraightFlush()
	assert.Equal(t, false, sf.IsFullHouse())

	fk := GenerateFourOfAKind()
	assert.Equal(t, false, fk.IsFullHouse())

	fh := GenerateFixFullHouse()
	assert.Equal(t, true, fh.IsFullHouse())

	f := GenerateFlush()
	assert.Equal(t, false, f.IsFullHouse())

	s := GenerateFixSuitStraight()
	assert.Equal(t, false, s.IsFullHouse())
}

func TestStraightFlushDeck_CompareTo(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()

	assert.Equal(t, 0, sf.CompareTo(sf))
	assert.Equal(t, 1, sf.CompareTo(fk))
	assert.Equal(t, 1, sf.CompareTo(fh))
	assert.Equal(t, 1, sf.CompareTo(f))
	assert.Equal(t, 1, sf.CompareTo(s))
}

func TestFourOfAKindDeck_CompareTo(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()

	assert.Equal(t, 2, fk.CompareTo(sf))
	assert.Equal(t, 0, fk.CompareTo(fk))
	assert.Equal(t, 1, fk.CompareTo(fh))
	assert.Equal(t, 1, fk.CompareTo(f))
	assert.Equal(t, 1, fk.CompareTo(s))
}

func TestFullHousDeck_CompareTo(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()

	assert.Equal(t, 2, fh.CompareTo(sf))
	assert.Equal(t, 2, fh.CompareTo(fk))
	assert.Equal(t, 0, fh.CompareTo(fh))
	assert.Equal(t, 1, fh.CompareTo(f))
	assert.Equal(t, 1, fh.CompareTo(s))
}

func TestUnmatchDeck_CompareTo(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFixFullHouse()
	f := GenerateFlush()
	s := GenerateFixSuitStraight()
	unm := Deck{
		Cards: []Card{
			Card{0, 1},
			Card{0, 2},
			Card{2, 13},
			Card{0, 5},
			Card{0, 6},
		},
	}

	assert.Equal(t, 2, unm.CompareTo(sf))
	assert.Equal(t, 2, unm.CompareTo(fk))
	assert.Equal(t, 2, unm.CompareTo(fh))
	assert.Equal(t, 2, unm.CompareTo(f))
	assert.Equal(t, 2, unm.CompareTo(s))
}
