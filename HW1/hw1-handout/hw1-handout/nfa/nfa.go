package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

// You may define helper functions here.

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO: Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	if len(input) <= 0 {
		return false
	}
	firstRune := input[0]

	canReach := false      // Set deafult return value to false
	currentState := start  // Set the start state as the current state
	var nextStates []state // Create a temp container to store next states
	var stateQueue []state // Create a queue to store current states
	queueLength := 0       // Set default queue length
	nextStates := TransitionFunction(currentState, firstRune)
	for stateIndex := 0; stateIndex < len(states); stateIndex++ {
		stateQueue = append(stateQueue, states[stateIndex])
	}
	nextStates = nil // remove data in temp container
	queueLength := len(stateQueue)

	for i := 1; i < len(input); i++ {
		currentRune := input[i]
		for queueIndex := 0; queueIndex < queueLength; queueIndex++ {
			temp := stateQueue[0] // first element in queue

		}
	}
}

// Pseudo - 21:24 version
// start 0
// abababa
// a: [1,2] queue1
// queue1.dequeu at many times as possble
// got the next states and save to queue2

// b: [0]  queue2
// queue2.dequeue as ,,,,,
// got the next states and save to queue1
// ....
// a: [1,2]
// b: [0]
// a: [1,2]
// b: [0]
// a: [1,2]
// for the non empty queue (queue1 or queue2)
// dequeu and see if match the target

// Pseudo - 21:37
// start 0
// abababa
// a: [1, 2] queue length = 2
// for i < queue length
// 		 temp = transitionFunction queue.poll()
// 		enque list
// queue = [0]
// b: [0] queue length=1
// again
// [1,2] queue length = 2
// ...

// Pseudo afternoon version
// queue to store final state node

// for each char in sequence:
// 	nextStateAmount = determine how many states we can reach from current state
// 	if nextStateAmount = 0:
// 		enqueue current state
// 	if nextStateAmount = 1;
// 		current state = next state
// 	if nextStateAmount > 1;
// 		recursioncall(next state1)
// 		recursioncall(next state2)
//         ...
//         recursioncall(next stateX)

// while queue is not empty
// 	poll queue one by one
// 	if equals to target state
// 		return true;
// 	else false
