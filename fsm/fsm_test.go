package fsm

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

var transitions = []Transition{
	{From: "Locked", Event: "Coin", To: "Unlocked", Action: "check"},
	//{From: "Locked", Event: "Push", To: "Locked", Action: "invalid-push"},
	{From: "Unlocked", Event: "Push", To: "Locked", Action: "pass"},
	{From: "Unlocked", Event: "Coin", To: "Unlocked", Action: "repeat-check"},
}

type TestActor struct {
	ActorDelegate
	fsm *StateMachine
}

func (actor *TestActor) init() {
	actor.InitState("Locked")
	actor.fsm = NewStateMachine(actor, transitions...)
}

func TestFSM(t *testing.T) {
	c.Convey("测试FSM", t, func() {
		actor := &TestActor{}
		actor.init()
		c.So(actor.fsm.Trigger("Push"), c.ShouldBeError)
		c.So(actor.CurrentState(), c.ShouldEqual, "Locked")
		c.So(actor.fsm.Trigger("Coin"), c.ShouldBeNil)
		c.So(actor.CurrentState(), c.ShouldEqual, "Unlocked")
		c.So(actor.fsm.Trigger("Coin"), c.ShouldBeNil)
		c.So(actor.CurrentState(), c.ShouldEqual, "Unlocked")
		c.So(actor.fsm.Trigger("Push"), c.ShouldBeNil)
		c.So(actor.CurrentState(), c.ShouldEqual, "Locked")
		c.So(actor.fsm.Trigger("Nothing"), c.ShouldBeError)
	})
}
