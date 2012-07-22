(declare primes)

(defn prime? [n]
  (let [factors (take-while #(<= % (Math/sqrt n)) primes)]
    (not-any? #(zero? (mod n %)) factors)))

(def primes
  (lazy-cat [2]
    (filter prime? (take-nth 2 (iterate inc 3)))))

(println (reduce + (take-while #(< % 2000000) primes)))
