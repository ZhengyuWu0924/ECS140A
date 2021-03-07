split([H|T], H, T).

not_reachable(_, State, DifferentState, []) :- State \= DifferentState.
not_reachable(Nfa, StartState, FinalState, Input) :-
    split(Input, Symbol, RestInput),
    transition(Nfa, StartState, Symbol, StateList),
    not_reachable_loop(Nfa, StateList, FinalState, RestInput)
.

not_reachable_loop(_, [], _, _) :- true.
not_reachable_loop(Nfa, StateList, FinalState, Input) :-
    split(StateList, FirstState, RestStateList),
    not_reachable(Nfa, FirstState, FinalState, Input),
    not_reachable_loop(Nfa, RestStateList, FinalState, Input)
.

reachable(Nfa, StartState, FinalState, Input) :-
    not(not_reachable(Nfa, StartState, FinalState, Input))
.