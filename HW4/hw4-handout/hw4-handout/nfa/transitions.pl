/* fooTransitions */
% 0 -a-> 1
% 0 -a-> 2
% 1 -b-> 3
% 2 -c-> 3
transition(fooTransitions, 0, a, [1,2]).
transition(fooTransitions, 0, b, []).
transition(fooTransitions, 0, c, []).
transition(fooTransitions, 1, a, []).
transition(fooTransitions, 1, b, [3]).
transition(fooTransitions, 1, c, []).
transition(fooTransitions, 2, a, []).
transition(fooTransitions, 2, b, []).
transition(fooTransitions, 2, c, [3]).

/* expTransitions */
% 0 -a-> 1
% 0 -a-> 2
% 0 -b-> 2
% 1 -b-> 0
transition(expTransitions, 0, a, [1,2]).
transition(expTransitions, 0, b, [2]).
transition(expTransitions, 1, a, []).
transition(expTransitions, 1, b, [0]).
transition(expTransitions, 2, a, []).
transition(expTransitions, 2, b, []).

/* langTransitions */
% 0 -a-> 0
% 0 -b-> 1
% 1 -a-> 1
% 1 -b-> 0
transition(langTransitions, 0, a, [0]).
transition(langTransitions, 0, b, [1]).
transition(langTransitions, 1, a, [1]).
transition(langTransitions, 1, b, [0]).

/* myTransitions */
% 0 -a -> 1
% 0 -a -> 2
% 0 -b -> 3
% 0 -c -> 4
% 1 -b -> 3
% 1 -b -> 4
% 2 -b -> 3
% 3 -c -> 3
% 3 -c -> 4
transition(myTransitions, 0, a, [1, 2]).
transition(myTransitions, 0, b, [3]).
transition(myTransitions, 0, c, [4]).
transition(myTransitions, 1, a, []).
transition(myTransitions, 1, b, [3, 4]).
transition(myTransitions, 1, c, []).
transition(myTransitions, 2, a, []).
transition(myTransitions, 2, b, [3]).
transition(myTransitions, 2, c, []).
transition(myTransitions, 3, a, []).
transition(myTransitions, 3, b, []).
transition(myTransitions, 3, c, [3, 4]).
transition(myTransitions, 4, a, []).
transition(myTransitions, 4, b, []).
transition(myTransitions, 4, c, []).