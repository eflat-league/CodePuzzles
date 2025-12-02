/*
Each node only transitions forward or ends the execution.
*/
package state_machine

import "github.com/eflat-league/CodePuzzles/leet_code/regex_evaluator/operators"

type Transition int

const (
	FORWARD Transition = iota
	END
	REGEX_MATCHED
)

type State struct {
	Input string
}

type StateMachineNode interface {
	Evaluate(*State) Transition
}

type RegexNode struct {
	RegexOperand operators.RegexOperand
}

// Returns true if regex matched, false if not
func ExecuteStateMachine(state State, nodes []StateMachineNode) bool {
	for _, n := range nodes {
		t := n.Evaluate(&state)
		switch t {
		case REGEX_MATCHED:
			return true
		case END:
			return false
		case FORWARD:
			continue
		default:
			panic("Unknown transition state")
		}
	}
	panic("Invalid state machine situation")
}

func (r *RegexNode) Evaluate(state *State) Transition {
	ok, result := r.RegexOperand.Evaluate(state.Input)
	if !ok {
		return END
	}
	state.Input = result
	return FORWARD
}

type EndNode struct{}

func (e *EndNode) Evaluate(state *State) Transition {
	//if all the previous nodes ate the entire string, success!
	if state.Input == "" {
		return REGEX_MATCHED
	}
	return END
}
