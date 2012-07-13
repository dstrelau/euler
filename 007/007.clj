(defn prime [n]
  (loop [p [2] i 3]
    (cond
      (= (count p) n) (last p)
      (some #(zero? (mod i %)) (take-while #(<= % (Math/sqrt i)) p)) (recur p (+ i 2))
      :else (recur (conj p i) (+ i 2)))))

(println (prime 10001))

; OR

(defn prime? [n, primes]
  (not-any? #(zero? (mod n %)) primes))

(def primes
  (lazy-cat [2 3] (filter #(prime? % primes) (take-nth 2 (iterate inc 5)))))

(println (nth primes 10000))
