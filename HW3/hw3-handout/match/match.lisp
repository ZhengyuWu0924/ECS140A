; You may define helper functions here
(defun match-recursion (flag pattern assertion)
  (if (equal flag t) 
    (cond
    ((and (null pattern) (null assertion)) t)
    ((null pattern) (match-recursion t pattern (cdr assertion)))
    ((null assertion) nil)
    ((or (equal (car pattern) '?) (equal (car pattern) '!))  (match-recursion t (cdr pattern) (cdr assertion)))
    ((equal (car pattern ) (car assertion)) (match-recursion nil (cdr pattern) (cdr assertion)))
    (t (match-recursion t pattern (cdr assertion)))
    )
    (cond
    ((and (null pattern) (null assertion)) t)
    ((or (null pattern) (null assertion)) nil)
    ((equal (car pattern) '?) (match-recursion nil (cdr pattern) (cdr assertion)))
    ((equal (car pattern) '!) (match-recursion t (cdr pattern) (cdr assertion)))
    (t (if (equal (car pattern ) (car assertion)) (match-recursion nil (cdr pattern) (cdr assertion)) nil))
    )
  )
)

(defun match (pattern assertion)
  (match-recursion nil pattern assertion)
)
