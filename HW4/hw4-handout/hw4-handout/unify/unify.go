package unify

import (
	"errors"
 	// "hw4/disjointset"
	"hw4/term"
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
	class map[*term.Term]*term.Term
	size map[*term.Term]int
	schema map[*term.Term]*term.Term
	visited map[*term.Term]bool
	acyclic map[*term.Term]bool
	vars map[*term.Term][]*term.Term
	result UnifyResult
}

// NewUnifier creates a struct of a type that satisfies the Unifier interface.
func NewUnifier() Unifier {
	panic("TODO: implement NewParser")
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

func (u *UnifierImpl) dfsInitialization(s *term.Term, t *term.Term) {

}

func (u *UnifierImpl) unifClosure(s *term.Term, t *term.Term) error {
	return nil
}

func (u *UnifierImpl) findSolution(s *term.Term) error {
	return nil
}

func (u *UnifierImpl) union(s *term.Term, t *term.Term) {

}

func (u *UnifierImpl) find(s *term.Term) {

}