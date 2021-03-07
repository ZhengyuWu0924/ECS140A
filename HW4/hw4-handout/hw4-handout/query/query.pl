# main :-
    # consult(["facts.pl"]).

/* All novels published either during the year 1953 or during the year 1996*/
check(Book) :- novel(Book, 1953).
check(Book) :- novel(Book, 1996).
year_1953_1996_novels(Book) :- check(Book).

/* List of all novels published during the period 1800 to 1900 (not inclusive)*/
period_1800_1900_novels(Book) :- novel(Book, Year), Year > 1800, Year < 1900.

/* Characters who are fans of LOTR */
mymembers(X, [X|_]). 
mymembers(X, [_|Y]) :- mymembers(X,Y).
lotr_fans(Fan) :- fan(Fan, Books), mymembers(the_lord_of_the_rings, Books).

/* Authors of the novels that heckles is fan of. */
# getname([],[]).
getname(Author, [H|_]) :- author(Author, Books), mymembers(H, Books).
getname(Author, [_|T]) :- getname(Author, T).
heckles_idols(Author) :- fan(heckles, List), getname(Author, List).

/* Characters who are fans of any of Robert Heinlein's novels */
chkfans(Fan, [H|_]) :- fan(Fan, List), mymembers(H, List).
chkfans(Fan, [_|T]) :- chkfans(Fan, T).
heinlein_fans(Fan) :- author(robert_heinlein, List), chkfans(Fan, List).

/* Novels common between either of Phoebe, Ross, and Monica */
chkmember(Book, [H|_], List) :- mymembers(H, List), Book = H.
chkmember(Book, [_|T], List) :- chkmember(Book, T, List).
mutual_novels(Book) :- fan(phoebe, List1), fan(monica, List3), chkmember(Book, List1, List3).
mutual_novels(Book) :- fan(ross, List2), fan(monica, List3), chkmember(Book, List2, List3).
mutual_novels(Book) :- fan(ross, List2), fan(phoebe, List1), chkmember(Book, List2, List1).
