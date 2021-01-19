package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

var res bool

// You may define helper functions here.

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// FINISHED: Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	res = false
	if len(input) == 0 { // Input nothing
		if start == final { // Corner case: start state = target state
			return true // Reachable since not moving
		}
		return false
	}
	backtrack(transitions, start, final, input)
	return res

	// var nextStates []state                            // Create a temp container to store next states
	// var stateQueue []state                            // Create a queue to store current states
	// transitionType := TransitionFunction(transitions) // New a TransitionFunction type
	// canReach := false                                 // Set deafult return value to false
	// currentState := start                             // Set the start state as the current state
	// firstRune := input[0]                             // Get the first moving command
	// queueLength := 0                                  // Set default queue length

	// // Get the possibility states by calling transition function
	// // with the current state and the command character
	// nextStates = transitionType(currentState, firstRune)

	// // Append all possible states to stateQueue
	// for stateIndex := 0; stateIndex < len(nextStates); stateIndex++ {
	// 	stateQueue = append(stateQueue, nextStates[stateIndex])
	// }
	// nextStates = nil // remove data in temp container

	// // For each command character in the input sequence
	// for i := 1; i < len(input); i++ {
	// 	currentRune := input[i]       // get the current character
	// 	queueLength = len(stateQueue) // get the queue length
	// 	// Use the character as transition command for every states in the queue
	// 	for queueIndex := 0; queueIndex < queueLength; queueIndex++ {
	// 		temp := stateQueue[0]                          // Pick first element from queue
	// 		stateQueue = stateQueue[1:]                    // Poll the first element from the queue
	// 		nextStates = transitionType(temp, currentRune) // Find its possible transitions
	// 		if len(nextStates) == 0 {                      // Current state has no transition to other states
	// 			nextStates = nil // Reset and empty the possible transitions container
	// 		} else { // Current staet has transitions to other states
	// 			// For every transable states
	// 			for stateIndex := 0; stateIndex < len(nextStates); stateIndex++ {
	// 				// Push it into the queue
	// 				stateQueue = append(stateQueue, nextStates[stateIndex])
	// 			}
	// 			nextStates = nil // Reset and empty the possible transitions container
	// 		}
	// 	}
	// 	// queueLength = len(stateQueue) // Update the length of the queue
	// }

	// for finalIndex := 0; finalIndex < len(stateQueue); finalIndex++ {
	// 	// If the target state is reachable after all command
	// 	if stateQueue[finalIndex] == final {
	// 		canReach = true
	// 	}
	// }
	// return canReach

}

func backtrack(
	trans TransitionFunction,
	start, final state,
	input []rune,
) {
	getTrans := TransitionFunction(trans)
	if len(input) == 0 {
		if start == final {
			res = true
		}
	} else {
		firstRune := input[0]
		input = input[1:]

		nextStates := getTrans(start, firstRune)
		if len(nextStates) == 0 {

		} else {
			for i := 0; i < len(nextStates); i++ {
				backtrack(trans, nextStates[i], final, input)
			}
		}

	}

}

// Pseudo - 2021.1.17 21:37
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
