:- initialization main.

main :-
    consult(['transitions.pl', 'nfa.pl']),
    (show_coverage(run_tests) ; true),
    halt.

:- begin_tests(nfa).

test(nfaExp1, [nondet]) :- reachable(expTransitions, 0, 2, [a]).
test(nfaExp2, [nondet]) :- reachable(expTransitions, 0, 2, [b]).
test(nfaExp3, [nondet]) :- reachable(expTransitions, 0, 1, [a, b, a]).
test(nfaExp4, [fail])   :- reachable(expTransitions, 0, 1, [a, b, a, b]).
test(nfaExp5, [nondet]) :- reachable(expTransitions, 0, 2, [a, b, a]).
test(nfaExp6, [nondet]) :- reachable(expTransitions, 0, 1, [a]).
test(nfaExp7, [fail]) :- reachable(expTransitions, 0, 1, [b]).
test(nfaExp8, [fail]) :- reachable(expTransitions, 0, 0, [a]).
test(nfaExp9, [fail]) :- reachable(expTransitions, 0, 0, [b]).
test(nfaExp10, [nondet]) :- reachable(expTransitions, 0, 0, [a, b]).
test(nfaExp11, [fail]) :- reachable(expTransitions, 0, 0, [a, b, a]).
test(nfaExp12, [fail]) :- reachable(expTransitions, 0, 2, [a, a, b]).
test(nfaExp13, [nondet]) :- reachable(expTransitions, 0, 2, [a, b, a, b, b]).


test(nfaFoo1, [nondet]) :- reachable(fooTransitions, 0, 3, [a, b]).
test(nfaFoo2, [nondet]) :- reachable(fooTransitions, 0, 3, [a, c]).
test(nfaFoo3, [nondet]) :- reachable(fooTransitions, 0, 1, [a]).
test(nfaFoo4, [fail])   :- reachable(fooTransitions, 0, 3, [a, a]).
test(nfaFoo5, [fail])   :- reachable(fooTransitions, 0, 3, [a]).
test(nfaFoo6, [fail])   :- reachable(fooTransitions, 0, 1, [b]).
test(nfaFoo7, [fail])   :- reachable(fooTransitions, 0, 1, [a, b]).
test(nfaFoo8, [fail]) :- reachable(fooTransitions, 0, 1, [a, c]).
test(nfaFoo9, [nondet]) :- reachable(fooTransitions, 0, 2, [a, b, a]).
test(nfaFoo10,[fail])   :- reachable(fooTransitions, 0, 2, [a, c]).
test(nfaFoo11, [nondet]):- reachable(fooTransitions, 0, 3, [a, b, a]).
test(nfaFoo12, [nondet]):- reachable(fooTransitions, 0, 3, [a, c, a, c]).

test(nfaLang1, [nondet]) :- reachable(langTransitions, 0, 0, [a, b, b]).
test(nfaLang2, [nondet]) :- reachable(langTransitions, 0, 1, [a, a, b]).
test(nfaLang3, [nondet]) :- reachable(langTransitions, 0, 0, [a, a, a, a, a]).
test(nfaLang4, [fail])   :- reachable(langTransitions, 0, 1, [a, a]).
test(nfaLang5, [fail])   :- reachable(langTransitions, 0, 0, [a, b, a, a]).
test(nfaLang6, [fail])   :- reachable(langTransitions, 0, 0, [b, b, b]).
test(nfaLang7, [nondet])   :- reachable(langTransitions, 0, 0, [b, b, b, b]).
test(nfaLang8, [fail])   :- reachable(langTransitions, 0, 1, [a, b, b, a]).
test(nfaLang9, [fail])   :- reachable(langTransitions, 0, 1, [a, b, a, a, b]).
test(nfaLang10, [nondet])   :- reachable(langTransitions, 0, 1, [a, b, a, b, a, b]).
test(nfaLang11, [nondet])   :- reachable(langTransitions, 1, 1, [a, a, a]).
test(nfaLang12, [fail])   :- reachable(langTransitions, 1, 1, [a, b, a]).
test(nfaLang13, [fail])   :- reachable(langTransitions, 1, 0, [a, a, a]).
test(nfaLang14, [nondet])   :- reachable(langTransitions, 1, 0, [a, b, a]).

test(nfaMy1, [nondet]) :- reachable(myTransitions, 0, 1, [a]).
test(nfaMy2, [nondet]) :- reachable(myTransitions, 0, 2, [a]).
test(nfaMy3, [fail]) :- reachable(myTransitions, 0, 3, [a]).
test(nfaMy4, [nondet]) :- reachable(myTransitions, 0, 3, [b]).
test(nfaMy5, [fail]) :- reachable(myTransitions, 0, 1, [a, a, b]).
test(nfaMy6, [nondet]) :- reachable(myTransitions, 0, 4, [c]).
test(nfaMy7, [nondet]) :- reachable(myTransitions, 0, 4, [a, b, c]).
test(nfaMy8, [nondet]) :- reachable(myTransitions, 0, 3, [a, b, c]).
test(nfaMy9, [fail]) :- reachable(myTransitions, 0, 2, [a, b, c]).

:- end_tests(nfa).