% A list is a 1-D array of numbers.
% A matrix is a 2-D array of numbers, stored in row-major order.

% You may define helper functions here.

% are_adjacent helper functions:
match_first_element([H|_], H).

are_adjacent_recursion([A|T], A, B) :-
    match_first_element(T, B)
.

are_adjacent_recursion([B|T], A, B) :-
    match_first_element(T, A)
.

are_adjacent_recursion([_|T], A, B) :-
    are_adjacent_recursion(T, A, B)
.

% matrix_transpose helper functions:
matrix_transpose_recursion([], []).
matrix_transpose_recursion([[]|_], []).
matrix_transpose_recursion(Matrix, Answer) :-
    construct_row(Matrix, MatrixNoFirstColumn, FirstAnswerRow),
    matrix_transpose_recursion(MatrixNoFirstColumn, SubAnswer),
    append_to_matrix(FirstAnswerRow, SubAnswer, Answer)
.

construct_row([], [], []).
construct_row([H|T], MatrixNoFirstColumn, Row) :-
    first_element(H, First),
    rest_elements(H, Rest),
    construct_row(T, SubMatrixNoFirstColumn, SubRow),
    append_to_matrix(Rest, SubMatrixNoFirstColumn, MatrixNoFirstColumn),
    append(First, SubRow, Row)
.

first_element([H|_], [H|[]]).

rest_elements([_|T], T).

append_to_matrix(Row, M, [Row|M]).

% are_neighbors helper functions:
are_not_neighbors(Matrix, A, B) :-
    matrix_transpose(Matrix, Answer),
    are_not_neighbors_recursion(Matrix, A, B),
    are_not_neighbors_recursion(Answer, A, B)
.

are_not_neighbors_recursion([], _, _) :- true.
are_not_neighbors_recursion([H|T], A, B) :-
    not(are_adjacent(H, A, B)),
    are_not_neighbors_recursion(T, A, B)
.

% are_adjacent(List, A, B) returns true iff A and B are neighbors in List.
are_adjacent(List, A, B) :-
    are_adjacent_recursion(List, A, B)
.

% matrix_transpose(Matrix, Answer) returns true iff Answer is the transpose of
% the 2D matrix Matrix.
matrix_transpose(Matrix, Answer) :-
    matrix_transpose_recursion(Matrix, Answer)
.

% are_neighbors(Matrix, A, B) returns true iff A and B are neighbors in the 2D
% matrix Matrix.
are_neighbors(Matrix, A, B) :-
    not(are_not_neighbors(Matrix, A, B))
.