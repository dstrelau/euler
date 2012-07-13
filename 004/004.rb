max = 0
999.downto(100) {|i|
  999.downto(100) {|j|
    n = i*j
    max = [max,n].max if n.to_s == n.to_s.reverse
  }
}
puts max
