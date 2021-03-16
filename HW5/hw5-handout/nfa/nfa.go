package nfa

import (
	"sync"
)

// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	var (
		sema = make(chan struct{}, 1)
		answer bool
	)
	reachable_concurrent(transitions, start, final, input, sema, &answer)
	return answer
}

func reachable_concurrent(
	transitions TransitionFunction,
	start, final state,
	input []rune,
	sema chan struct{},
	answer *bool,
) {
	if len(input) == 0 {
		if start == final {
			sema <- struct{}{}
			*answer = true
			<-sema
		}
	} else {
		var wg sync.WaitGroup
		for _, next := range transitions(start, input[0]) {
			wg.Add(1)
			go func(next state) {
				defer wg.Done()
				reachable_concurrent(transitions, next, final, input[1:], sema, answer)
			}(next)
			wg.Wait()
			if *answer {
				return
			}
		}
	}
}