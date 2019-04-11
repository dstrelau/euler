package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type solution struct {
	Number   int
	Answer   int
	Duration time.Duration
}

type solutionSlice []solution

func (s solutionSlice) Len() int           { return len(s) }
func (s solutionSlice) Less(i, j int) bool { return s[i].Number < s[j].Number }
func (s solutionSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	problems := make(map[int]func() int)

	if len(os.Args) > 1 {
		// if we got args, solve those problems only
		for _, arg := range os.Args[1:] {
			n, _ := strconv.Atoi(arg)
			if solution, solved := solvers[n]; solved {
				problems[n] = solution
			}
		}
	} else {
		// if we didn't get args, solve everything
		problems = solvers
	}

	// work on all solvers in parallel
	var solutions []solution
	var wg sync.WaitGroup
	out := make(chan solution)

	for n, solver := range problems {
		wg.Add(1)
		go func(n int, f func() int) {
			t0 := time.Now()
			answer := f()
			duration := time.Now().Sub(t0)
			out <- solution{Number: n, Answer: answer, Duration: duration}
		}(n, solver)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for solution := range out {
		solutions = append(solutions, solution)
		wg.Done()
	}

	sort.Sort(solutionSlice(solutions))
	for _, s := range solutions {
		fmt.Printf("%03d: %v (%v)\n", s.Number, s.Answer, s.Duration)
	}
}

var solvers = map[int]func() int{
	// 002: Find the sum of even-valued fibonacci numbers below 4M
	2: func() int {
		done := make(chan bool)

		fib := fibGen(done)
		evenFib := filterInt(fib, func(n int) bool { return n%2 == 0 })
		evenBelow4M := takeUntil(evenFib, done, func(n int) bool { return n > 4000000 })
		return sumIntC(evenBelow4M)
	},

	// 003: What is the largest prime factor of 600851475143?
	3: func() int { return maxInt(primeFactors(600851475143)) },

	// 004: Find the largest palindrome made from the product of two 3-digit numbers.
	4: func() int {
		max := 0
		for i := 999; i > 99; i-- {
			for j := 999; j > 99; j-- {
				product := i * j
				forwards := strconv.Itoa(product)
				backwards := reverse(forwards)
				if forwards == backwards {
					max = maxI(max, product)
				}
			}
		}
		return max
	},

	// 005: Find the smallest positive number evenly divisible by all of (1..20)
	// 006: Find the difference of sum of squares and square of the sum of (1..100)
	// 007: What is the 10001st prime number?
	7: func() int {
		primes := nPrimes(10001)
		return primes[len(primes)-1]
	},
	// 008: Find the greatest product of five consecutive digits in a 1000-digit number.
	// 009: Find the product abc for the pythagorean triple where a + b + c = 1000.

	// 010: Find the sum of all the primes below two million.
	10: func() int {
		primes := primesUpto(2000000 - 1)
		return sumSlice(primes)
	},

	// 011: Find the largest product of four adjacent numbers in a grid
	// 012: What is the first triangle number to have over 500 divisors?
	// 013: What are the first ten digits of the sum of 100 50-digit numbers?

	// 017: How many letters are used writing out numbers one to one thousand?
	17: func() int {
		var buffer bytes.Buffer
		lookupTable := buildIntWordTable()
		for i := 1; i <= 1000; i++ {
			eye := numInWords(lookupTable, i)
			for _, word := range strings.Split(eye, " ") {
				buffer.WriteString(word)
			}
		}
		return len(buffer.String())
	},

	// 018: Find the maximum sum traversing a number triangle top to bottom
	18: func() int {
		triangle := readIntMatrix("data/018")

		for i := len(triangle) - 1; i > 0; i-- {
			for j := 0; j < len(triangle[i-1]); j++ {
				a, b := triangle[i][j], triangle[i][j+1]
				triangle[i-1][j] += maxI(a, b)
			}
		}

		return triangle[0][0]
	},

	// 019: How many Sundays fell on the 1st of the month from 1901-2000?
	19: func() int {
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
	},

	// 020: Find the sum of the digits of 100!
	20: func() int {
		n := big.NewInt(0).MulRange(1, 100)
		sum := 0
		for _, m := range n.String() {
			o, _ := strconv.Atoi(string(m))
			sum += o
		}
		return sum
	},

	// 021: Evaluate the sum of all the amicable numbers under 10000.
	21: func() int {
		result := 0
		toSum := make(map[int]struct{})
		sumOfDivisors := make(map[int]int)
		for i := 0; i <= 10000; i++ {
			sumOfDivisors[i] = sumSlice(divisors(i))
		}
		for a := 0; a <= 10000; a++ {
			b := sumOfDivisors[a]
			if sumOfDivisors[b] == a && a != b {
				toSum[a] = struct{}{}
				toSum[b] = struct{}{}
			}
		}
		for n := range toSum {
			result += n
		}
		return result
	},

	// 022: Compute the sum of character-position scores for a word list
	22: func() int {
		a := int('A')
		names := strings.Split(readFile("data/022"), ",")
		for i, n := range names {
			names[i] = strings.Trim(n, "\"")
		}
		sort.Strings(names)

		result := 0
		for i, n := range names {
			score := 0
			for _, c := range n {
				score += int(c) - a + 1
			}
			result += score * (i + 1)
		}
		return result
	},
}

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
