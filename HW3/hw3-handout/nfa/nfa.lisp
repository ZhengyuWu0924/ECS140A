(defun recursive-calls (transition state-list final input)
    (cond
    ((null state-list) nil)
    (t (or (reachable transition (car state-list) final input)
        (recursive-calls transition (cdr state-list) final input))))
)

(defun reachable (transition start final input)
    (cond
    ((null input) (if (equal start final) t nil))
    (t (recursive-calls transition (funcall transition start (car input)) final (cdr input))))
)