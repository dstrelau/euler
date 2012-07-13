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

p prime_factors(600851475143).max
p max_prime_factor(600851475143)
