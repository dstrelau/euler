def nth_prime(n)
  primes = [2]
  i = 3
  while primes.length < n
    primes << i unless primes.any? {|p| i % p == 0 }
    i += 2
  end
  primes[-1]
end

puts nth_prime(10_001)
