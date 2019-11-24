package main

import (
	"fmt"
	"github.com/looplab/fsm"
	"math/rand"
	"time"
	//"reflect"
	"unsafe"
)

type Door struct {
	To  string
	FSM *fsm.FSM
	//EFSM *fsm.EFSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}
	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed", "half open"}, Dst: "open"},
			{Name: "close", Src: []string{"open", "half open"}, Dst: "closed"},
			{Name: "half", Src: []string{"open", "closed"}, Dst: "half open"},
			{Name: "break", Src: []string{"open", "closed", "half open"}, Dst: "broken"},
			{Name: "all", Src: []string{"open", "closed", "half open", "broken"}, Dst: "all"},
		},
		fsm.Callbacks{
			"enter_state":  func(e *fsm.Event) { d.enterState(e) },
			"after_half":   func(e *fsm.Event) { d.enterHalfState(e) },
			"before_event": func(e *fsm.Event) { d.beforeEvent(e) },
		},
	)
	return d
}

var target_event string

func (d *Door) beforeEvent(e *fsm.Event) {
	fmt.Println("beforeevent", e.Event, "src", e.Src, "dst", e.Dst, "args", e.Args)
	if e.Event == "all" {
		e.Cancel()
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(4)
		fmt.Println(a)
		switch a {
		case 0:
			target_event = "open"
		case 1:
			target_event = "close"
		case 2:
			target_event = "half"
		case 3:
			target_event = "break"
		default:
			target_event = ""
		}
		return
	}
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Println(d.To, e.FSM.Current(), e.Src, e.Dst, e.Args)
}

func (d *Door) enterHalfState(e *fsm.Event) {
	fmt.Println("now door is half")
}

func main() {
	door := NewDoor("heaven")
	door.FSM.Event("half", 102, 3123)
	// you can only transition once
	door.FSM.Event("all")
	if len(target_event) == 0 {
		return
	}
	fmt.Println("te: ", target_event)
	door.FSM.Event(target_event)

	door.FSM.Event("half")
	door.FSM.Event("half")
	door.FSM.Event("open")
	door.FSM.Event("close")
	door.FSM.Event("open")
	door.FSM.Event("half")
	//door.FSM.Event("break")
	fmt.Println(door.FSM.AvailableTransitions())
	fmt.Println(door.FSM.Current())
	fmt.Println(unsafe.Sizeof(*door.FSM))
}
