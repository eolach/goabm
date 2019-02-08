package abm

import "sync"

type ABM struct {
}

func (a *ABM) AddAgent(ag Agent) {

}

func (a *ABM) StartSimulation() {
	for i := 0; i < a.Limit(); i++ {
		a.World().Tick()
	}
	var wg sync.WaitGroup
	for j := 0; j < a.AgentCount(); j++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, j int) {
			a.agents[j].Run()
			wg.Done()
		}(&wg, j)
	}
	wg.Wait()
	if a.reportFunc != nil {
		a.reportFunc(a)
	}
}

// Agent defines independent behaviour
type Agent interface {
	Run()
}

// World represents space in which agend dwell and
// interact. It update the state on each Tick()
type World interface {
	Tick() // mark the beginning of the next time period
}
