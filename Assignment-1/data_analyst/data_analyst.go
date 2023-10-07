package data_analyst

type DataAnalyst struct {
	position string
	salary   float64
	address  string
}

func (p *DataAnalyst) GetPosition() string {
	return p.position
}

func (p *DataAnalyst) GetSalary() float64 {
	return p.salary
}

func (p *DataAnalyst) GetAddress() string {
	return p.address
}

func (p *DataAnalyst) SetPosition(pos string) {
	p.position = pos
}

func (p *DataAnalyst) SetSalary(sal float64) {
	p.salary = sal
}

func (p *DataAnalyst) SetAddress(adr string) {
	p.address = adr
}
