package mrctech

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateStraightFlush(t *testing.T) {
	for i := 0; i < 5; i++ {
		GenerateStraightFlush().DisplayDeck()
	}
}

func TestGenerateFourOfAKind(t *testing.T) {
	for i := 0; i < 5; i++ {
		GenerateFourOfAKind().DisplayDeck()
	}
}

func TestGenerateFullHouse(t *testing.T) {
	for i := 0; i < 5; i++ {
		GenerateFullHouse().DisplayDeck()
	}
}

func TestGenerateFlush(t *testing.T) {
	for i := 0; i < 5; i++ {
		GenerateFlush().DisplayDeck()
	}
}

func TestGenerateStraight(t *testing.T) {
	for i := 0; i < 5; i++ {
		GenerateStraight().DisplayDeck()
	}
}

func TestDeck_IsFlush(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFullHouse()
	f := GenerateFlush()
	s := GenerateStraight()

	assert.Equal(t, true, sf.IsFlush())
	assert.Equal(t, false, fk.IsFlush())
	assert.Equal(t, false, fh.IsFlush())
	assert.Equal(t, true, f.IsFlush())
	assert.Equal(t, false, s.IsFlush())
}

// Flush 會產生同花順 有時會false暫時忽略
func TestDeck_IsStraight(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFullHouse()
	f := GenerateFlush()
	s := GenerateStraight()

	assert.Equal(t, true, sf.IsStraight())
	assert.Equal(t, false, fk.IsStraight())
	assert.Equal(t, false, fh.IsStraight())
	assert.Equal(t, false, f.IsStraight())
	assert.Equal(t, true, s.IsStraight())
}

func TestDeck_IsStraightFlush(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFullHouse()
	f := GenerateFlush()
	s := GenerateStraight()

	assert.Equal(t, true, sf.IsStraightFlush())
	assert.Equal(t, false, fk.IsStraightFlush())
	assert.Equal(t, false, fh.IsStraightFlush())
	assert.Equal(t, false, f.IsStraightFlush())
	assert.Equal(t, false, s.IsStraightFlush())
}

func TestDeck_IsFourOfAKind(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFullHouse()
	f := GenerateFlush()
	s := GenerateStraight()

	assert.Equal(t, false, sf.IsFourOfAKind())
	assert.Equal(t, true, fk.IsFourOfAKind())
	assert.Equal(t, false, fh.IsFourOfAKind())
	assert.Equal(t, false, f.IsFourOfAKind())
	assert.Equal(t, false, s.IsFourOfAKind())
}

func TestDeck_IsFullHouse(t *testing.T) {
	sf := GenerateStraightFlush()
	fk := GenerateFourOfAKind()
	fh := GenerateFullHouse()
	f := GenerateFlush()
	s := GenerateStraight()

	assert.Equal(t, false, sf.IsFullHouse())
	assert.Equal(t, false, fk.IsFullHouse())
	assert.Equal(t, true, fh.IsFullHouse())
	assert.Equal(t, false, f.IsFullHouse())
	assert.Equal(t, false, s.IsFullHouse())
}

func TestDeck_CompareTo(t *testing.T) {
	d1 := GenerateStraightFlush()
	d2 := GenerateStraightFlush()
	assert.Equal(t, 0, d1.CompareTo(d2))
}
