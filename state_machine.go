package jot

type StateMachine interface {
	State() string
	Transition(state string, data interface{}) error
}

type NullStateMachine struct {
}

func (sm *NullStateMachine) State() string {
	return ""
}

func (sm *NullStateMachine) Transition(state string, data interface{}) error {
	return nil
}
