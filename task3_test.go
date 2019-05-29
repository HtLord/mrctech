package mrctech

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestFibonacci(t *testing.T) {
	assert.Equal(t, 0, Fibonacci(0))
	assert.Equal(t, 1, Fibonacci(1))
	assert.Equal(t, 1, Fibonacci(2))
	assert.Equal(t, 2, Fibonacci(3))
	assert.Equal(t, 3, Fibonacci(4))
	assert.Equal(t, 5, Fibonacci(5))
	assert.Equal(t, 8, Fibonacci(6))
	assert.Equal(t, 13, Fibonacci(7))
	assert.Equal(t, 144, Fibonacci(12))
}

func TestFibonacci2(t *testing.T) {
	t.Log(Fibonacci(0))
	t.Log(Fibonacci(1))
	t.Log(Fibonacci(2))
	t.Log(Fibonacci(3))
	t.Log(Fibonacci(4))
	t.Log(Fibonacci(5))
	t.Log(Fibonacci(6))
	t.Log(Fibonacci(7))
	t.Log(Fibonacci(12))
}