package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func intGen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func takeUntil(in <-chan int, done chan<- bool, f func(int) bool) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			if f(n) {
				close(done)
				close(out)
				return
			}
			out <- n
		}
	}()

	return out
}

func filterInt(in <-chan int, f func(int) bool) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			if f(n) {
				out <- n
			}
		}
		close(out)
	}()

	return out
}

func fibGen(done <-chan bool) chan int {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = j, i+j {
			select {
			case <-done:
				close(c)
				return
			case c <- j:
			}
		}
	}()

	return c
}

func multiplesOfBelowLimit(multiples []int, limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		for _, m := range multiples {
			if i%m == 0 {
				sum += i
			}
		}
	}
	return sum
}

// sieve of Eratosthenes
// false values represent primes
func primesUpto(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	a := make([]bool, limit+1)

	for i := 2; float64(i) < math.Sqrt(float64(limit)); i++ {
		if !a[i] {
			for j := i * i; j <= limit; j += i {
				a[j] = true
			}
		}
	}

	primes := []int{}
	for i := 2; i < len(a); i++ {
		if !a[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

// Prime Number Theorem
func nPrimes(n int) []int {
	if n < 6 {
		return primesUpto(11)[:n]
	}

	n64 := float64(n)
	logN := math.Log(n64)
	limit := n64*logN + n64*math.Log(logN)
	return primesUpto(int(limit))[:n]
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
	var rv []int
	for k := range factors {
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

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(data), "\n")
}

func readIntMatrix(filename string) [][]int {
	lines := strings.Split(readFile(filename), "\n")
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

func maxInt(a []int) int {
	max := 0
	for _, i := range a {
		max = maxI(max, i)
	}
	return max
}

func maxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func daysIn(year int, month int) int {
	switch month {
	case 4, 6, 9, 11:
		return 30
	case 2:
		if year%4 == 0 {
			if year%100 == 0 && year%400 != 0 {
				return 28
			}
			return 29
		}
		return 28
	default:
		return 31
	}
}

func sumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

func sumIntC(in <-chan int) int {
	sum := 0
	for n := range in {
		sum += n
	}
	return sum
}

func divisors(n int) []int {
	r := []int{1}
	root := int(math.Sqrt(float64(n)))
	for i := 2; i <= root; i++ {
		if n%i == 0 {
			r = append(r, i)
			if i != n/i {
				r = append(r, n/i)
			}
		}
	}
	return r
}
