package computer_vision_engineer

type CVEngineer struct {
	position string
	salary   float64
	address  string
}

func (p *CVEngineer) GetPosition() string {
	return p.position
}

func (p *CVEngineer) GetSalary() float64 {
	return p.salary
}

func (p *CVEngineer) GetAddress() string {
	return p.address
}

func (p *CVEngineer) SetPosition(pos string) {
	p.position = pos
}

func (p *CVEngineer) SetSalary(sal float64) {
	p.salary = sal
}

func (p *CVEngineer) SetAddress(adr string) {
	p.address = adr
}
