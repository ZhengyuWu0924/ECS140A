isFinalState(CurrentState, FinalState) :- CurrentState == FinalState.
isEmpty([]).
inputIsEmpty([]).

nextSym([H|_], H).
nextSym([_|T], NextSym) :- nextSym(T, NextSym).

nextState([H|_], H]).
nextState([_|T], NextState) :- nextState(T, nextState).

nextStateInList([H|_], NextOne) :- NextOne = H.
nextStateInList([_|T], NextOne) :- nextStateInList(T, NextOne).
popFirst([_|T], Poped) :- Poped = T.

innerStateCheck([H|_], Target) :- H is Target.
innerStateCheck([_|T], Target) :- innerStateCheck(T, Target).

isEmptyTransition(Nfa, State, Sym) :- transition(Nfa, State, Sym, List), isEmpty(List).
isInnerTransition(Nfa, State, Sym) :- transition(Nfa, State, Sym, List), innerStateCheck(List, State).

noOutterTransition(Nfa, State, [H|_]) :- isEmptyTransition(Nfa, State, H).
noOutterTransition(Nfa, State, [_|T]) :- noOutterTransition(Nfa, State, T).
noOutterTransition(Nfa, State, [H|_]) :- isInnerTransition(Nfa, State, H).
noOutterTransition(Nfa, State, [_|T]) :- noOutterTransition(Nfa, State, T).


reachable(Nfa, StartState, FinalState, Input) :- 
    inputIsEmpty(Input),
    isFinalState(StartState, FinalState).
reachable(Nfa, StartState, FinalState, Input) :-
    nextSym(Input, NextSymbol),
    noOutterTransition(Nfa, StartState, NextSymbol),
    isFinalState(StartState, FinalState).
reachable(Nfa, StartState, FinalState, Input) :-
    nextSym(Input, NextSymbol),
    transition(Nfa, StartState, NextSymbol, StateList), 
    reachable(Nfa, nextStateInList(StateList, NextOnes), FinalState, popFirst(Input, Res)).

reachable(Nfa, StaretState, FinalState, Input) :-
    nextSym(Input, NextSymbol),
    popFirst(Input, Res),
    inputIsEmpty(Res),
    transition(Nfa, StartState, NextSymbol, StateList),
    nextStateInList(StateList, NextOnes),
    isFinalState(NextOnes, FinalState).
reachable(Nfa, StaretState, FinalState, Input) :-
    nextSym(Input, NextSymbol),
    transition(Nfa, StartState, NextSymbol, StateList),
    nextStateInList(StateList, NextOnes),
    isFinalState(NextOnes, FinalState).