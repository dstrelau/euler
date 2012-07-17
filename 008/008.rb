N = File.readlines('number.txt').join

puts N.each_char.each_cons(5).inject(0) {|max,arr|
  [max, arr.map(&:to_i).inject(&:*)].max
}
