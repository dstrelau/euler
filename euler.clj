(require '(clojure [string :as s]))

(defmacro euler [n body] `(defmethod solution ~n [n#] (delay ~body)))
(defmulti solution (fn [n] n))
(defmethod solution :default [n] (delay "UNSOLVED"))
(defn solve [ens]
  (doseq [n ens
        :let [start (System/nanoTime)
              sol @(solution n)
              stop (System/nanoTime)
              elapsed (float (/ (- stop start) 1000000000))]]
    (println (format "%03d: %s (%.3fs)" n sol elapsed))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defn multiples-of [divisors]
  (filter
    (fn [n] (some #(zero? (mod n %)) divisors))
    (rest (range))))

(def fib-seq (lazy-cat [1 2]
  (map + fib-seq (rest fib-seq))))

(defn prime-factors [num]
  (loop [n num f 2 factors []]
    (cond
      (= n 1) (distinct factors)
      (zero? (mod n f)) (recur (quot n f) f (conj factors f))
      :else (recur n (inc f) factors))))

(defn palindrome? [n]
  (= (str n) (s/reverse (str n))))

(defn gcd [a b] (loop [a a b b] (if (zero? b) a (recur b (mod a b)))))
(defn lcm [a b] (/ (* a b) (gcd a b)))

(defn square [n] (* n n))
(def squares (lazy-seq (map square (rest (range)))))

(defn sum-of-squares [n] (reduce + (take n squares)))
(defn square-of-sum [n] (square (reduce + (range (inc n)))))

(defn nth-prime [n]
  (loop [p [2] i 3]
    (cond
      (= (count p) n) (last p)
      (some #(zero? (mod i %)) (take-while #(<= % (Math/sqrt i)) p)) (recur p (+ i 2))
      :else (recur (conj p i) (+ i 2)))))

(declare primes)

(defn prime? [n]
  (let [factors (take-while #(<= % (Math/sqrt n)) primes)]
    (not-any? #(zero? (mod n %)) factors)))

(def primes
  (lazy-cat [2]
    (filter prime? (take-nth 2 (iterate inc 3)))))

(defn readlines [filename] (s/split-lines (slurp filename)))

; partition then reduce over the partitions
(defn part-reduce [size reduce-fn nums]
  (let [parts (partition size 1 nums)]
    (map #(reduce reduce-fn %) parts)))

(defn transpose [m]
  (apply vector (apply map vector m)))

(defn diagonals [m]
  (let [n (count m)
        at (fn [i j] (get (get m i) j))]
    (for [diag (range (dec (* 2 n)))
          :let [startrow (max 0 (inc (- diag n)))]]
      (for [col (range startrow (inc (- diag startrow)))
            :let [row (- diag col)]]
        (at row col)))))

(defn diagonalsr [m]
  (-> m reverse vec diagonals))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

; 001: Find the sum of all the multiples of 3 or 5 below 1000.
(euler 1 (reduce + (take 1000 (multiples-of [3,5]))))

; 002: Find the sum of even-valued fibonacci numbers below 4M
(euler 2 (reduce + (filter even? (take-while #(< % 4000000) fib-seq))))

; 003: What is the largest prime factor of 600851475143?
(euler 3 (reduce max (prime-factors 600851475143)))

; 004: Find the largest palindrome made from the product of two 3-digit numbers.
(euler 4
  (apply max
    (filter palindrome?
      (for [i (range 100 999) j (range 100 999)] (* i j)))))

; 005: Find the smallest positive number evenly divisible by all of (1..20)
(euler 5 (reduce lcm (range 1 20)))

; 006: Find the difference of sum of squares and square of the sum of (1..100)
(euler 6
  (let [n 100] (- (square-of-sum n) (sum-of-squares n))))

; 007: What is the 10001st prime number?
(euler 7 (nth primes 10000))

; 008: Find the greatest product of five consecutive digits in a 1000-digit number.
(euler 8
  (let [n (map #(Integer. (str %))
               (.replaceAll (slurp "data/008") "[\r\n]+" ""))]
    (reduce max (part-reduce 5 * n))))

; 009: Find the product abc for the pythagorean triple where a + b + c = 1000.
(euler 9 (first
  (for [n (range 1 1000) m (range 1 n)
        :let [a (- (* n n) (* m m))
              b (* 2 n m)
              c (+ (* n n) (* m m))]
        :when (= 1000 (+ a b c))]
    (* a b c))))

; 010: Find the sum of all the primes below two million.
(euler 10 (reduce + (take-while #(< % 2000000) primes)))

; 011: Find the largest product of four adjacent numbers in a grid
(euler 11
  (let [toi #(Integer. %)
      break #(s/split % #"\s+")
      data (readlines "data/011")
      grid (mapv #(mapv toi (break %)) data)
      maxr #(reduce max (flatten %))
      adjmax (fn [m] (maxr (remove empty? (map #(part-reduce 4 * %) m))))]
    (reduce max (map adjmax
      [grid (transpose grid) (diagonals grid) (diagonalsr grid)]))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;
(let [args (seq *command-line-args*)]
  (solve (map #(Integer. %) (or args (range 1 11)))))
;;;;;;;;;;;;;;;;;;;;;;;;;;;
