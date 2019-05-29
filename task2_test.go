package mrctech

import (
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T){
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<10; i++{
		suit := rand.Intn(4)
		t.Log(suit)
	}
}