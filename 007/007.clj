(defn prime [n]
  (loop [p [2] i 3]
    (cond
      (= (count p) n) (last p)
      (some #(zero? (mod i %)) (take-while #(<= % (Math/sqrt i)) p)) (recur p (+ i 2))
      :else (recur (conj p i) (+ i 2)))))

(println (prime 10001))

; OR

(declare primes)

(defn prime? [n]
  (let [factors (take-while #(<= % (Math/sqrt n)) primes)]
    (not-any? #(zero? (mod n %)) factors)))

(def primes
  (lazy-cat [2]
    (filter prime? (take-nth 2 (iterate inc 3)))))

(println (nth primes 10000))
