(println (first
  (for [n (range 1 1000) m (range 1 n)
        :let [a (- (* n n) (* m m))
              b (* 2 n m)
              c (+ (* n n) (* m m))]
        :when (= 1000 (+ a b c))]
    (* a b c))))
