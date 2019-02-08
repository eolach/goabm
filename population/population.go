package main

import (
	"github.com/eolach/goabm/abm"
	"github.com/eolach/goabm/models/human"
	"github.com/eolach/goabm/ui/term"
)

// Agent defines independent behaviour
type Agent interface {
	Run()
}

// World represents space in which agend dwell and
// interact. It update the state on each Tick()
type World interface {
	Tick() // mark the beginning of the next time period
}

// Main function
func main() {
	a := abm.New()
	a.SetWorld(grid.New(100, 100))
	a.AddAgent(&Human{})

	ch := make(chan int)
	a.SetReportFunc(func(a *abm.ABM) {
		ch <- a.Count(func(agent abm.Agent) bool {
			return agent(*human.Human).IsAlive()
		})
	})
	go a.StartSimulation()

	ui := term.NewUI()
	ui.AddChart("Humans Alive", ch)
	ui.Loop()
}
