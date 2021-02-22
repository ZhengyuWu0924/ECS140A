; A list is a 1-D array of numbers.
; A matrix is a 2-D array of numbers, stored in row-major order.

; If needed, you may define helper functions here.
(defun are-neighbers-recursion (matrix a b)
    (cond
    ((or (null matrix) (null (car matrix))) nil)
    (t (or (are-adjacent (car matrix) a b) (are-neighbers-recursion (cdr matrix) a b)))
    )
)
; AreAdjacent returns true iff a and b are adjacent in lst.
(defun are-adjacent (lst a b)
    (cond
    ((or (null lst) (null (cdr lst))) nil)
    ((equal (car lst) a) (if (equal (car (cdr lst)) b) t (are-adjacent (cdr (cdr lst)) a b)))
    ((equal (car lst) b) (if (equal (car (cdr lst)) a) t (are-adjacent (cdr (cdr lst)) a b)))
    (t (are-adjacent (cdr lst) a b))
    )
)

; Transpose returns the transpose of the 2D matrix mat.
(defun transpose (matrix)
    (cond
    ((or (null matrix) (null (car matrix))) nil)
    (t (cons (mapcar #'car matrix) (transpose (mapcar #'cdr matrix))))
    )
)

; AreNeighbors returns true iff a and b are neighbors in the 2D
; matrix mat.
(defun are-neighbors (matrix a b)
    (or (are-neighbers-recursion matrix a b) (are-neighbers-recursion (transpose matrix) a b))
)

;; (write (are-neighbors '( (1 2 3) ) 1 2))
