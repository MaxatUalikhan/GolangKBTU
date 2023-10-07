package frontend_developer

type FrontendDeveloper struct {
	position string
	salary   float64
	address  string
}

func (p *FrontendDeveloper) GetPosition() string {
	return p.position
}

func (p *FrontendDeveloper) GetSalary() float64 {
	return p.salary
}

func (p *FrontendDeveloper) GetAddress() string {
	return p.address
}

func (p *FrontendDeveloper) SetPosition(pos string) {
	p.position = pos
}

func (p *FrontendDeveloper) SetSalary(sal float64) {
	p.salary = sal
}

func (p *FrontendDeveloper) SetAddress(adr string) {
	p.address = adr
}
