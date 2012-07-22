def solution(n, &block)
  @solutions ||= Hash.new {|h,k| h[k] = [] }
  @solutions[n] << block
end

def solve(ens)
  ens = ens.size == 0 ? @solutions.keys : ens.map(&:to_i)
  ens.each do |n|
    puts "NOT SOLVED" if @solutions[n].size == 0
    @solutions[n].each do |b|
      t = Time.now
      puts "#{n.to_s.rjust(3,'0')}: #{b.call} (#{(Time.now - t).round(3)}s)"
    end
  end
end

#############################################################################

def fib_seq(max)
  a = [1,2]
  while a[-1] < max
    a << a[-1] + a[-2]
  end
  a[0...-1]
end

def prime_factors(n)
  factors = Hash.new(0)
  f = 2
  while n > 1
    if n % f == 0
      n /= f
      factors[f] += 1
    else
      f += 1
    end
  end
  factors.keys
end

# golf'd
def max_prime_factor(n)
  f = 2
  n % f == 0 ? n /= f : f += 1 while n > 1
  f
end

def nth_prime(n)
  primes = [2]
  i = 3
  while primes.length < n
    unless primes.any? {|p| break if p > Math.sqrt(i) ; i % p == 0 }
      primes << i
    end
    i += 2
  end
  primes[-1]
end

Primes = Enumerator.new do |yielder|
  yielder << 2
  primes = [2]
  i = 3
  loop do
    unless primes.any? {|p| break if p > Math.sqrt(i) ; i % p == 0 }
      primes << i
      yielder << i
    end
    i += 2
  end
end

# start with fac = n! and try to divide by i = (2..n),
# storing fac=(fac/i) if fac/i is still divisible by all (1..n)
def smallest_divisible_under(n)
  fac = (1..n).inject(&:*)
  i = 2
  while i <= n
    if (1..n).all? {|j| (fac/i) % j == 0 }
      fac /= i
    else
      i += 1
    end
  end
  fac
end

def gcd(a,b)
  if b == 0
  then a
  else gcd(b, a % b)
  end
end

def lcm(a,b)
  (a*b) / gcd(a,b)
end

def sum_of_squares(n)
  (1..n).map {|i| i**2 }.inject(&:+)
end

def square_of_sum(n)
  (1..n).inject(&:+)**2
end

# method so we can break out of both loops. </3 ruby
def triple(sum)
  (1...sum).each {|a|
    (1...sum).each {|b|
      c = Math.sqrt(a*a + b*b)
      return (a*b*c) if a + b + c == sum
    }
  }
end

# faster: pythagorean triples have a formula!
def triple2(sum)
  (1..sum).each {|m|
    (1..sum).each {|n|
      a = n*n - m*m
      b = 2*n*m
      c = n*n + m*m
      return (a*b*c) if a + b + c == sum
    }
  }
end

#############################################################################

### PROBLEM 001 ###
# Find the sum of all the multiples of 3 or 5 below 1000.
solution(1) {
  (1...1000).select {|n| n % 3 == 0 || n % 5 == 0 }.inject(&:+)
}

### PROBLEM 002 ###
# Find the sum of even-valued fibonacci numbers below 4M
solution(2) {
  fib_seq(4_000_000).select(&:even?).inject(&:+)
}

### PROBLEM 003 ###
# What is the largest prime factor of 600851475143?
solution(3) { prime_factors(600851475143).max }
solution(3) { max_prime_factor(600851475143) }

### PROBLEM 004 ###
# Find the largest palindrome made from the product of two 3-digit numbers.
solution(4) do
  max = 0
  999.downto(100) {|i|
    999.downto(100) {|j|
      n = i*j
      max = [max,n].max if n.to_s == n.to_s.reverse
    }
  }
end

### PROBLEM 005 ###
# Find the smallest positive number evenly divisible by all of (1..20)
solution(5) { smallest_divisible_under(20) }
solution(5) { (1..20).inject(1) {|acc,n| lcm(acc,n) } }

### PROBLEM 006 ###
# Difference of sum of squares and square of sum for (1..100)
solution(6) { square_of_sum(100) - sum_of_squares(100) }

### PROBLEM 007 ###
# What is the 10001st prime number?
solution(7) { nth_prime(10_001) }

### PROBLEM 008 ###
# Find the greatest product of five consecutive digits in the 1000-digit number.
solution(8) {
  number = File.readlines('008_number.txt').join
  number.each_char.each_cons(5).inject(0) {|max,arr|
    [max, arr.map(&:to_i).inject(&:*)].max
  }
}

### PROBLEM 009 ###
# Find the product abc for the pythagorean triple where a + b + c = 1000.
solution(9) { triple(1000) }
solution(9) { triple2(1000) }

### PROBLEM 010 ###
# Find the sum of all the primes below two million.
solution(10) { Primes.take_while {|n| n < 2_000_000 }.inject(&:+) }

#############################################################################

solve(ARGV) if __FILE__ == $0
