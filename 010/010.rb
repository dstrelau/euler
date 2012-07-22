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

puts Primes.take_while {|n| n < 2_000_000 }.inject(&:+)
