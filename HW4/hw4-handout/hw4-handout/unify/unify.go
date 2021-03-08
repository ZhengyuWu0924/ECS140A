package unify

import (
	"errors"
	// "hw4/disjointset"
	"hw4/term"
	//"fmt"
)

// ErrUnifier is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrUnifier = errors.New("unifier error")

// UnifyResult is the result of unification. For example, for a variable term
// `s`, `UnifyResult[s]` is the term which `s` is unified with.
type UnifyResult map[*term.Term]*term.Term

// Unifier is the interface for the term unifier.
// Do not change the definition of this interface
type Unifier interface {
	Unify(*term.Term, *term.Term) (UnifyResult, error)
}

type UnifierImpl struct {
	class   map[*term.Term]*term.Term
	size    map[*term.Term]int
	schema  map[*term.Term]*term.Term
	visited map[*term.Term]bool
	acyclic map[*term.Term]bool
	vars    map[*term.Term][]*term.Term
	result  UnifyResult
}

// NewUnifier creates a struct of a type that satisfies the Unifier interface.
func NewUnifier() Unifier {
	return &UnifierImpl{
		class:   make(map[*term.Term]*term.Term),
		size:    make(map[*term.Term]int),
		schema:  make(map[*term.Term]*term.Term),
		visited: make(map[*term.Term]bool),
		acyclic: make(map[*term.Term]bool),
		vars:    make(map[*term.Term][]*term.Term),
		result:  make(UnifyResult),
	}
}

func (u *UnifierImpl) Unify(s *term.Term, t *term.Term) (UnifyResult, error) {
	err := u.unifClosure(s, t)
	if err != nil {
		return nil, err
	}
	err = u.findSolution(s)
	if err != nil {
		return nil, err
	}
	return u.result, nil
}

func (u *UnifierImpl) unifClosure(s *term.Term, t *term.Term) error {
	u.init(s)
	u.init(t)
	s, t = u.find(s), u.find(t)
	if s != t {
		ss, st := u.schema[s], u.schema[t]
		if ss.Typ != term.TermVariable && st.Typ != term.TermVariable {
			var ssLiteral, stLiteral string
			var ssArgsLength, stArgsLength int = 0, 0

			if ss.Typ == term.TermCompound {
				ssLiteral = ss.Functor.Literal
				ssArgsLength = len(ss.Args)
			} else {
				ssLiteral = ss.Literal
			}
			if st.Typ == term.TermCompound {
				stLiteral = st.Functor.Literal
				stArgsLength = len(st.Args)
			} else {
				stLiteral = st.Literal
			}

			// check if the number of arguments are equal for schema(s) and schema(t)
			// if they are compound terms.
			if ssArgsLength != stArgsLength {
				return ErrUnifier
			}

			if ssLiteral == stLiteral {
				u.union(s, t)
				if ssArgsLength > 0 {
					for i := range ss.Args {
						if err := u.unifClosure(ss.Args[i], st.Args[i]); err != nil {
							return ErrUnifier
						}
					}
				}
			} else {
				return ErrUnifier
			}
		} else {
			u.union(s, t)
		}
	}
	return nil
}

func (u *UnifierImpl) findSolution(s *term.Term) error {
	u.init(s)
	s = u.schema[u.find(s)]
	if u.acyclic[s] {
		return nil
	}
	if u.visited[s] {
		return ErrUnifier
	}
	if s.Typ == term.TermCompound {
		u.visited[s] = true
		for _, arg := range s.Args {
			if err := u.findSolution(arg); err != nil {
				return ErrUnifier
			}
		}
		u.visited[s] = false
	}
	u.acyclic[s] = true
	for _, x := range u.vars[u.find(s)] {
		if x != s {
			u.result[x] = s
		}
	}
	return nil
}

func (u *UnifierImpl) union(s *term.Term, t *term.Term) {
	sizeS, sizeT := u.size[s], u.size[t]
	if sizeS >= sizeT {
		u.size[s] = sizeS + sizeT
		u.vars[s] = append(u.vars[s], u.vars[t]...)
		if u.schema[s].Typ == term.TermVariable {
			u.schema[s] = u.schema[t]
		}
		u.class[t] = s
	} else {
		u.size[t] = sizeT + sizeS
		u.vars[t] = append(u.vars[t], u.vars[s]...)
		if u.schema[t].Typ == term.TermVariable {
			u.schema[t] = u.schema[s]
		}
		u.class[s] = t
	}
}

func (u *UnifierImpl) find(s *term.Term) *term.Term {
	var t *term.Term
	if u.class[s] == s {
		return s
	}
	t = u.find(u.class[s])
	u.class[s] = t
	return t
}

func (u *UnifierImpl) init(s *term.Term) {
	if _, ok := u.class[s]; !ok {
		u.class[s] = s
		u.schema[s] = s
		u.size[s] = 1
		u.visited[s] = false
		u.acyclic[s] = false
		if s.Typ == term.TermVariable {
			u.vars[s] = []*term.Term{s}
		} else {
			u.vars[s] = []*term.Term{}
		}
	}
}
