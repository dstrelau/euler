package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntGen(t *testing.T) {
	out := intGen(2)
	assert.Equal(t, 1, <-out)
	assert.Equal(t, 2, <-out)
	_, more := <-out
	assert.False(t, more)
}

func TestFilterInt(t *testing.T) {
	in := intGen(4)

	out := filterInt(in, func(n int) bool { return n%2 == 0 })
	assert.Equal(t, 2, <-out)
	assert.Equal(t, 4, <-out)
	assert.Equal(t, 0, <-out)

}

func TestFibGen(t *testing.T) {
	done := make(chan bool)
	defer close(done)

	expected := []int{1, 1, 2, 3, 5, 8, 13}
	fib := fibGen(done)

	for _, n := range expected {
		assert.Equal(t, n, <-fib)
	}
}

func TestTakeUntil(t *testing.T) {
	seq := func(done <-chan bool) chan int {
		out := make(chan int)
		i := 0
		go func() {
			for {
				select {
				case <-done:
					close(out)
					return
				default:
					i++
					out <- i
				}
			}
		}()

		return out
	}

	done := make(chan bool)
	expected := []int{1, 2, 3, 4, 5}
	untilFive := func(n int) bool { return n > 5 }

	integers := seq(done)
	takeUntil5 := takeUntil(integers, done, untilFive)

	i := 0
	for n := range takeUntil5 {
		assert.Equal(t, expected[i], n)
		i++
	}
}

func TestSum(t *testing.T) {
	assert.Equal(t, 15, sumIntC(intGen(5)))
}

func TestPrimesUpto(t *testing.T) {
	assert.Equal(t, []int{}, primesUpto(1))
	assert.Equal(t, []int{2}, primesUpto(2))
	assert.Equal(t, []int{2, 3}, primesUpto(3))
	assert.Equal(t, []int{2, 3, 5, 7}, primesUpto(10))
	assert.Equal(t, []int{2, 3, 5, 7, 11, 13, 17, 19}, primesUpto(20))
}

func TestNPrimes(t *testing.T) {
	assert.Equal(t, []int{2}, nPrimes(1))
	assert.Equal(t, []int{2, 3, 5, 7, 11}, nPrimes(5))
	assert.Equal(t, []int{2, 3, 5, 7, 11, 13}, nPrimes(6))
}
