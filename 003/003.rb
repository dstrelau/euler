def prime_factors(n)
  factors = []
  f = 2
  while n > 1 && f < n
    if n % f == 0
      n /= f
      factors << f
    end
    f += 1
  end
  factors
end

# golf'd
def max_prime_factor(n)
  f = 2
  n % f == 0 ? n /= f : f += 1 while n > 1
  f
end

# p prime_factors(600851475143)
p max_prime_factor(600851475143)
