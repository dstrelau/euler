(defn square [n] (* n n))
(def squares (lazy-seq (map square (rest (range)))))

(defn sum-of-squares [n] (reduce + (take n squares)))
(defn square-of-sum [n] (square (reduce + (range (inc n)))))

(let [n 100]
  ; (println (square-of-sum n))
  ; (println (sum-of-squares n))
  (println (- (square-of-sum n) (sum-of-squares n))))
