package main

/*
mController :
	StateMachineCreate({
		"wait" : func() return WaitStateCreate(Entity, gMap),
		"move" : func() return MoveStateCreate(gHero, gMap),
	})
*/

//
// StateMachine - a state machine
//
// Usage:
//
// gStateMachine = StateMachineCreate(
// {
// 		"MainMenu" : func() State {
// 			return MainMenuCreate()
// 		},
// 		"InnerGame" : func() State {
// 			return InnerGameCreate()
// 		},
// 		"GameOver" : func() State {
// 			return GameOverCreate()
// 		},
// })
// gStateMachine.Change("MainGame")
//

//StateMachine mController
type StateMachine struct {
	states  map[string]func() State
	current State
}

func StateMachineCreate(states map[string]func() State) *StateMachine {
	return &StateMachine{
		states:  states,
		current: nil,
	}
}

//Change state
// e.g. mController.Change("move", {x = -1, y = 0})
func (m *StateMachine) Change(stateName string, enterParams Direction) {
	if m.current != nil {
		m.current.Exit()
	}
	m.current = m.states[stateName]()
	m.current.Enter(enterParams) //thinking.. pass enterParams
}

//Update(dt)
func (m *StateMachine) Update(dt float64) {
	m.current.Update(dt)
}

func (m *StateMachine) Render() {
	m.current.Render()
}
