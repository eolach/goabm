package human

// Human implements Agent for human that can age
type Human struct {
	age  int
	dead bool
}

func (h *Human) Run() {
	if h.dead {
		return
	}
	h.age++
	if h.age == AveDeathAge {
		h.Die()
	}
}

func (h *Human) Die() {
	h.dead = true
}
