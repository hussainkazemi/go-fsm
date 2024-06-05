package main

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"time"
)

func Red(tr *TrafficLight) {
	fmt.Println("Traffic light is Red, Pleas stop")
	time.Sleep(5 * time.Second)
	tr.FSM.Event(context.Background(), "Green")
}

func Yellow(tr *TrafficLight) {
	fmt.Println("Traffic light is Yellow, Pleas go fast")
	time.Sleep(2 * time.Second)
	tr.FSM.Event(context.Background(), "Red")
}

func Green(tr *TrafficLight) {
	fmt.Println("Traffic light is Green, you can move")
	time.Sleep(5 * time.Second)
	tr.FSM.Event(context.Background(), "Yellow")
}

type TrafficLight struct {
	name string
	FSM  *fsm.FSM
}

func (tr *TrafficLight) Start() {
	err := tr.FSM.Event(context.Background(), "Green")
	if err != nil {
		fmt.Errorf("an error occurred %w", err)
	}
}

func NewTrafficLight() *TrafficLight {
	tr := &TrafficLight{}
	tr.FSM = fsm.NewFSM(
		"Red",
		fsm.Events{
			{Name: "Red", Src: []string{"Yellow"}, Dst: "Red"},
			{Name: "Green", Src: []string{"Red"}, Dst: "Green"},
			{Name: "Yellow", Src: []string{"Green"}, Dst: "Yellow"},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { tr.enterState(e) },
		},
	)
	return tr
}

func (tr *TrafficLight) enterState(e *fsm.Event) {
	switch e.Dst {
	case "Red":
		Red(tr)
	case "Yellow":
		Yellow(tr)
	case "Green":
		Green(tr)
	}
}

func main() {
	tr := NewTrafficLight()
	tr.Start()
}
