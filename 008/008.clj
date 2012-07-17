(def N (map
         #(Integer. (str %))
         (.replaceAll (slurp "number.txt") "[\r\n]+" "")))

(println (reduce max (map #(reduce * %) (partition 5 1 N))))
