isEmpty([]).
get(Nfa, Start, Sym, Y) :- transition(Nfa, Start, Sym, Y), isEmpty(Y).