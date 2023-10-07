package hr

type HR struct {
	position string
	salary   float64
	address  string
}

func (p *HR) GetPosition() string {
	return p.position
}

func (p *HR) GetSalary() float64 {
	return p.salary
}

func (p *HR) GetAddress() string {
	return p.address
}

func (p *HR) SetPosition(pos string) {
	p.position = pos
}

func (p *HR) SetSalary(sal float64) {
	p.salary = sal
}

func (p *HR) SetAddress(adr string) {
	p.address = adr
}
