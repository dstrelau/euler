(defn prime-factors [num]
  (loop [n num f 2 factors []]
    (cond
      (= n 1) factors
      (> f n) factors
      (zero? (mod n f)) (recur (quot n f) (inc f) (conj factors f))
      :else (recur n (inc f) factors))))

; (println (prime-factors 600851475143))
(println (reduce max (prime-factors 600851475143)))
