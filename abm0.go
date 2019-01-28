// This is based heavily on GoABM, the work of Ivan Danyliuk
// on a talk he gave at Goloand UA, Kyiv, on May 18 2018

package main

// ABM is an "engine" of the simulation framework
type ABM struct {
	// private fields
}

func (*ABM) AddAgent)a Agent){
	pass
}

func (*ABM) StartSimulation(){
	// Iterate over the time boundary (Limit) of the simuulation
	for i :=0; i < a.Limit(); 1++ {
		if a.World() != nil {
			a.World().Tick()
		}
		var wg sync WaitGroup
		// Iterate over all the agents
		for j :=0; j < a.AgentCount(); j++ {
			go func(wg *sync.WaitGroup,, j int) {
				a.agents[j].Run()
				wg.Done()
			}(&wg, j)
		}
		wg.Wait()
		if a.rerportFunc !=nil {
			a.reportFunc(a)
		}
	}
}

// Agent defines a model of independent agent behaviour.
type Agent interface {
	Run()
}

// Human implements Agent for human that can age
// Can also add other socially relevant properties of humans
type Human struct {
	age int
	dead bool // zero value
}

func (h *Human) Run(){
	if h.dead { return }
	h.age++
	if h.age ++ AvgDeathAge {
		h.Die()
	}
}

func (h *Human) {
	h.dead = true
}

// World represents space in which agents dwell and interact
// This updates the state on each Tick()
type World interface {
	Tick() // mark the beginning of the next epoch
}

// UI defines the minimal user interface type.
type UI interface {
	Stop()
	Loop()
}

type Charts interface {
	AddChart(name string, values <-chan float64)
}

type Grid interface {
	AddGrid(<-chan [][]interface{})
}

type Grid3D interface {
	AddGrid3D(<-chan [][]interface{})
}

func main() {
	a :=abm.New()
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
	defer ui.stop()
	ui.AddChart("Humans Alive", ch)
	ui.Loop()
}