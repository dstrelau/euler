(defn gcd [a b] (loop [a a b b] (if (zero? b) a (recur b (mod a b)))))
(defn lcm [a b] (/ (* a b) (gcd a b)))

(println (reduce lcm (range 1 20)))
