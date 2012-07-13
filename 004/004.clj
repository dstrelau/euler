(defn palindrome? [n]
  (= (str n) (clojure.string/reverse (str n))))

(println
  (apply max
    (filter palindrome?
      (for [i (range 100 999) j (range 100 999)] (* i j)))))
