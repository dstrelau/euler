import Text.Printf
import Data.List
import Data.Maybe

-- returns true if f is a factor of n
divides :: Integer -> Integer -> Bool
divides f n = n `rem` f == 0

-- returns True if all of fs are factors of n
hasFactors :: [Integer] -> Integer -> Bool
hasFactors fs n = all (\f -> f n) (map divides fs)

-- a lazy list of fibonacci numbers
fibs :: [Integer]
fibs = 0 : 1 : zipWith (+) fibs (tail fibs)

-- returns the prime factors of n
primeFactors :: Integer -> [Integer]
primeFactors n = factorize 2 n where
  factorize _ 1 = []
  factorize d n
    | d * d > n = [n]
    | d `divides` n = d : factorize d (n `div` d)
    | otherwise = factorize (d + 1) n

-- returns primes
primes :: [Integer]
primes = 2 : filter (null . tail . primeFactors) [3,5..]

-- returns True if n is a palindrome number
isPalindrome :: Integer -> Bool
isPalindrome n = (show n) == (reverse $ show n)

-- returns pythagorean triples below
pythag :: [(Integer, Integer, Integer)]
pythag = [(a,b,c) | m <- [2..],
                    n <- [1..(m-1)],
                    let a = m^2 - n^2,
                    let b = 2*n*m,
                    let c = n^2 + m^2,
                    a^2 + b^2 == c^2]

solve :: Integer -> Integer
-- 001: Find the sum of all the multiples of 3 or 5 below 1000.
solve 1 = sum $ [n | n <- [1..999], 3 `divides` n || 5 `divides` n]
-- 002: Find the sum of even-valued fibonacci numbers below 4M
solve 2 = sum $ filter even $ takeWhile (< 4000000) fibs
-- 003: What is the largest prime factor of 600851475143?
solve 3 = last $ primeFactors 600851475143
-- 004: Find the largest palindrome made from the product of two 3-digit numbers.
solve 4 = maximum [ z | x <- [100..999], y <- [100..999], let z = x*y, isPalindrome z]
-- 005: Find the smallest positive number evenly divisible by all of (1..20)
solve 5 = foldl1 lcm [1..20]
-- 006: Find the difference of sum of squares and square of the sum of (1..100)
solve 6 = ((^2) $ sum [1..100]) - (sum [x^2 | x <- [1..100]])
-- 007: What is the 10001st prime number?
solve 7 = primes !! 10000
-- 008: Find the greatest product of five consecutive digits in a 1000-digit number.
-- 009: Find the product abc for the pythagorean triple where a + b + c = 1000.
solve 9 = let (a,b,c) = fromJust $ (find (\(a,b,c) -> a + b + c == 1000) pythag) in a*b*c
-- 010: Find the sum of all the primes below two million.
-- 011: Find the largest product of four adjacent numbers in a grid
-- 012: What is the first triangle number to have over 500 divisors?
-- 013: What are the first ten digits of the sum of 100 50-digit numbers?
-- 014: Find the starting number that produces the longest chain
solve n = -1

main = do
  mapM_ putStrLn [printf "%03d: %d" n (solve n) | n <- [1..15]]
