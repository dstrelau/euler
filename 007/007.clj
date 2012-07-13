(defn prime [n]
  (loop [p [2] i 3]
    (cond
      (= (count p) n) (last p)
      (some #(zero? (mod i %)) p) (recur p (+ i 2))
      :else (recur (conj p i) (+ i 2)))))

(println (prime 10001))
