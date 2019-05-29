package mrctech

import (
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
