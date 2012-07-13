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

puts smallest_divisible_under(20)
