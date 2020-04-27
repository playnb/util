package fsm

/*
	fork from https://github.com/smallnest/gofsm
*/
import "fmt"

type Transition struct {
	From   string
	Event  string
	To     string
	Action string
}

type StateMachine struct {
	delegate    Delegate
	transitions []Transition
}

type Error interface {
	error
	BadEvent() string
	CurrentState() string
}

type fsmError struct {
	badEvent     string
	currentState string
}

func (e fsmError) Error() string {
	return fmt.Sprintf("state machine error: cannot find transition for event [%s] when in state [%s]\n", e.badEvent, e.currentState)
}

func (e fsmError) BadEvent() string {
	return e.badEvent
}

func (e fsmError) CurrentState() string {
	return e.currentState
}

type Delegate interface {
	HandleEvent(action string, fromState string, toState string, args ...interface{}) error
	CurrentState() string
}

func NewStateMachine(delegate Delegate, transitions ...Transition) *StateMachine {
	return &StateMachine{delegate: delegate, transitions: transitions}
}

func (m *StateMachine) Trigger(event string, args ...interface{}) error {
	var err error
	current := m.delegate.CurrentState()
	trans := m.findTransMatching(current, event)
	if trans == nil {
		return fsmError{event, current}
	}
	err = m.delegate.HandleEvent(trans.Action, current, trans.To, args)
	return err
}

func (m *StateMachine) findTransMatching(fromState string, event string) *Transition {
	for _, v := range m.transitions {
		if v.From == fromState && v.Event == event {
			return &v
		}
	}
	return nil
}
