package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

type solution func() int

func main() {
	solutions := defineSolutions()
	nums := os.Args[1:]
	if len(nums) > 0 {
		for _, nStr := range nums {
			n, _ := strconv.Atoi(nStr)
			if solutions[n] != nil {
				solve(n, solutions[n])
			} else {
				fmt.Printf("%03d: NOT SOLVED")
			}
		}
	} else {
		for k, v := range solutions {
			solve(k, v)
		}
	}
}

func solve(n int, s solution) {
	t0 := time.Now()
	result := s()
	t1 := time.Now()
	fmt.Printf("%03d: %v (%v)\n", n, result, t1.Sub(t0))
}

/////////////////////////////////////////////////////////////////////////////

func fibGen() chan int {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = j, i+j {
			c <- j
		}
	}()

	return c
}

func primeFactors(n int) []int {
	factors := make(map[int]struct{})
	f := 2
	for n > 1 {
		if n%f == 0 {
			n /= f
			factors[f] = struct{}{}
		} else {
			f++
		}
	}
	rv := make([]int, 0)
	for k, _ := range factors {
		rv = append(rv, k)
	}
	return rv
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func buildIntWordTable() map[int]string {
	return map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety",
	}
}

func numInWords(lookupTable map[int]string, n int) string {
	result := []string{}

	thousands := n / 1000
	remainder := n - 1000*thousands
	if thousands > 0 {
		result = append(result, lookupTable[thousands], "thousand")
	}

	hundreds := remainder / 100
	remainder = remainder - 100*hundreds
	if hundreds > 0 {
		result = append(result, lookupTable[hundreds], "hundred")
		if remainder > 0 {
			result = append(result, "and")
		}
	}

	word, found := lookupTable[remainder]
	if found {
		result = append(result, word)
	} else {
		tens := remainder / 10
		result = append(result, lookupTable[10*tens], lookupTable[remainder%10])
	}

	return strings.Join(result, " ")
}

func readFileLines(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.Trim(string(data), "\n"), "\n")
}

func readIntMatrix(filename string) [][]int {
	lines := readFileLines(filename)
	matrix := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		nums := make([]int, len(line))
		for j, v := range line {
			n, _ := strconv.Atoi(v)
			nums[j] = n
		}
		matrix[i] = nums
	}
	return matrix
}

func maxOf(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func daysIn(year int, month int) int {
	switch month {
	case 4, 6, 9, 11:
		return 30
	case 2:
		if year%4 == 0 {
			if year%100 == 0 && year%400 != 0 {
				return 28
			} else {
				return 29
			}
		} else {
			return 28
		}
	default:
		return 31
	}
}

/////////////////////////////////////////////////////////////////////////////

func defineSolutions() map[int]solution {
	solutions := make(map[int]solution)

	// 001: Find the sum of all the multiples of 3 or 5 below 1000.
	solutions[1] = func() int {
		sum := 0
		for i := 0; i < 1000; i++ {
			if i%3 == 0 || i%5 == 0 {
				sum += i
			}
		}
		return sum
	}

	// 002: Find the sum of even-valued fibonacci numbers below 4M
	solutions[2] = func() int {
		fib := fibGen()
		sum := 0

		for n := 0; n < 4000000; {
			n = <-fib
			if n%2 == 0 {
				sum += n
			}
		}
		return sum
	}

	// 003: What is the largest prime factor of 600851475143?
	solutions[3] = func() int {
		max := 0
		for _, n := range primeFactors(600851475143) {
			max = maxOf(max, n)
		}
		return max
	}

	// 004: Find the largest palindrome made from the product of two 3-digit numbers.
	solutions[4] = func() int {
		max := 0
		for i := 999; i > 99; i-- {
			for j := 999; j > 99; j-- {
				product := i * j
				forwards := strconv.Itoa(product)
				backwards := reverse(forwards)
				if forwards == backwards {
					max = maxOf(max, product)
				}
			}
		}
		return max
	}

	// 005: Find the smallest positive number evenly divisible by all of (1..20)

	// 006: Find the difference of sum of squares and square of the sum of (1..100)
	// 007: What is the 10001st prime number?
	// 008: Find the greatest product of five consecutive digits in a 1000-digit number.
	// 009: Find the product abc for the pythagorean triple where a + b + c = 1000.
	// 010: Find the sum of all the primes below two million.
	// 011: Find the largest product of four adjacent numbers in a grid
	// 012: What is the first triangle number to have over 500 divisors?
	// 013: What are the first ten digits of the sum of 100 50-digit numbers?

	// 017: How many letters are used writing out numbers one to one thousand?
	solutions[17] = func() int {
		var buffer bytes.Buffer
		lookupTable := buildIntWordTable()
		for i := 1; i <= 1000; i++ {
			eye := numInWords(lookupTable, i)
			for _, word := range strings.Split(eye, " ") {
				buffer.WriteString(word)
			}
		}
		return len(buffer.String())
	}

	// 018: Find the maximum sum traversing a number triangle top to bottom
	solutions[18] = func() int {
		triangle := readIntMatrix("data/018")

		for i := len(triangle) - 1; i > 0; i-- {
			for j := 0; j < len(triangle[i-1]); j++ {
				a, b := triangle[i][j], triangle[i][j+1]
				triangle[i-1][j] += maxOf(a, b)
			}
		}

		return triangle[0][0]
	}

	// 019: How many Sundays fell on the 1st of the month from 1901-2000?
	solutions[19] = func() int {
		count := 0
		dayOfWeek := 2 // 1901-01-01 = Tues

		for year := 1901; year <= 2000; year++ {
			for month := 1; month <= 12; month++ {
				dayOfWeek = (dayOfWeek + daysIn(year, month)) % 7
				if dayOfWeek == 0 {
					count++
				}
			}
		}

		return count
	}

	// 020: Find the sum of the digits of 100!
	solutions[20] = func() int {
		n := big.NewInt(0).MulRange(1, 100)
		sum := 0
		for _, m := range n.String() {
			o, _ := strconv.Atoi(string(m))
			sum += o
		}
		return sum
	}

	return solutions
}
