package mrctech

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiniProduct(t *testing.T) {
	assert.Equal(t, 0, MiniProduct([]int{0, 1, 2, 3}))
	assert.Equal(t, -6, MiniProduct([]int{-1, 1, 2, 3}))
	assert.Equal(t, -12, MiniProduct([]int{-1, -2, 2, 3}))
	assert.Equal(t, -12, MiniProduct([]int{-1, -2, 0, 1, 2, 3}))
	assert.Equal(t, -1, MiniProduct([]int{-1, 0, 0, 0}))
	assert.Equal(t, -2, MiniProduct([]int{-1, -2, 0, 0, 0}))
	assert.Equal(t, -2, MiniProduct([]int{-1, 0, 0, 2}))
	assert.Equal(t, -4, MiniProduct([]int{-1, -2, 0, 0, 2}))
}

func TestMiniProduct2(t *testing.T) {
	t.Log(MiniProduct([]int{0, 1, 2, 3}))
	t.Log(MiniProduct([]int{-1, 1, 2, 3}))
	t.Log(MiniProduct([]int{-1, -2, 2, 3}))
	t.Log(MiniProduct([]int{-1, -2, 0, 1, 2, 3}))
	t.Log(MiniProduct([]int{-1, 0, 0, 0}))
	t.Log(MiniProduct([]int{-1, -2, 0, 0, 0}))
	t.Log(MiniProduct([]int{-1, 0, 0, 2}))
	t.Log(MiniProduct([]int{-1, -2, 0, 0, 2}))
}
