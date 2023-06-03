#lang racket

(define (is-alien-sorted words order)
  (let ([dict (gen-dict order)])
    (begin
      (define (compare words dict)
        (if (null? (cdr words))
            #t
            (if (smaller-word (car words) (car (cdr words)) dict)
                (compare (cdr words) dict)
                #f)))
      (compare words dict))))

(define (gen-dict order)
  (begin
    (define dict (make-hash))
    (for ([i (in-range (string-length order))])
      (hash-set! dict (string-ref order i) i))
    dict))

(define (smaller-word w1 w2 dict)
  (let ([l1 (string-length w1)] [l2 (string-length w2)])
    (cond
      [(and (= l1 0) (>= l2 0)) #t]
      [(and (> l1 0) (= l2 0)) #f]
      [else
       (let ([o1 (hash-ref dict (string-ref w1 0))]
             [o2 (hash-ref dict (string-ref w2 0))])
         (cond
           [(< o1 o2) #t]
           [(> o1 o2) #f]
           [(= o1 o2)
            (smaller-word (substring w1 1) (substring w2 1) dict)]))])))

;; #t
(printf "~a\n"
        (is-alien-sorted (list "hello" "leetcode")
                         "hlabcdefgijkmnopqrstuvwxyz"))
;; #f
(printf "~a\n"
        (is-alien-sorted (list "word" "world" "row")
                         "worldabcefghijkmnpqstuvxyz"))
;; #f
(printf "~a\n"
        (is-alien-sorted (list "apple" "app") "abcdefghijklmnopqrstuvwxyz"))
