def fib(max)
  a = [1,2]
  while a[-1] < max
    a << a[-1] + a[-2]
  end
  a[0...-1]
end

puts fib(4_000_000).select(&:even?).inject(&:+)
