package fsm

import "github.com/playnb/util"

type ActorDelegate struct {
	state           string
	changeTime      int64
	OnExit          func(fromState string, args ...interface{})
	Action          func(action string, fromState string, toState string, args ...interface{}) error
	OnActionFailure func(action string, fromState string, toState string, err error, args ...interface{}, )
	OnEnter         func(toState string, args []interface{})
}

func (actor *ActorDelegate) InitState(state string) {
	actor.state = state
}

func (actor *ActorDelegate) CurrentState() string {
	return actor.state
}

func (actor *ActorDelegate) CurrentTimestamp() int64 {
	return actor.changeTime
}

func (actor *ActorDelegate) HandleEvent(action string, fromState string, toState string, args ...interface{}) error {
	var err error
	if len(action) != 0 {
		if actor.Action != nil {
			err = actor.Action(action, fromState, toState, args)
		}
		if err != nil {
			if actor.OnActionFailure != nil {
				actor.OnActionFailure(action, fromState, toState, err, args)
			}
			return nil
		}
	}
	if fromState != toState {
		if actor.OnExit != nil {
			actor.OnExit(fromState, args)
		}
	}
	actor.state = toState
	if fromState != toState {
		actor.changeTime = util.NowTimestamp()
		if actor.OnEnter != nil {
			actor.OnEnter(toState, args)
		}
	}

	return nil
}
