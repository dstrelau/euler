# method so we can break out of both loops. </3 ruby
def triple(sum)
  (1...sum).each {|a|
    (1...sum).each {|b|
      c = Math.sqrt(a*a + b*b)
      return (a*b*c) if a + b + c == sum
    }
  }
end

puts triple(1000)

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

puts triple2(1000)
