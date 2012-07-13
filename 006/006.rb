def sum_of_squares(n)
  (1..n).map {|i| i**2 }.inject(&:+)
end

def square_of_sum(n)
  (1..n).inject(&:+)**2
end

i = 100
# puts sum_of_squares(i)
# puts square_of_sum(i)
puts square_of_sum(i) - sum_of_squares(i)
