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

puts nth_prime(10_001)
