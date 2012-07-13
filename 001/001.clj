(defn multiples-of [divisors]
  (filter
    (fn [n] (some #(zero? (mod n %)) divisors))
    (rest (range))))

(println reduce + (take 1000 (multiples-of [3,5])))
