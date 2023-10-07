package project_manager

type ProjectManager struct {
	position string
	salary   float64
	address  string
}

func (p *ProjectManager) GetPosition() string {
	return p.position
}

func (p *ProjectManager) GetSalary() float64 {
	return p.salary
}

func (p *ProjectManager) GetAddress() string {
	return p.address
}

func (p *ProjectManager) SetPosition(pos string) {
	p.position = pos
}

func (p *ProjectManager) SetSalary(sal float64) {
	p.salary = sal
}

func (p *ProjectManager) SetAddress(adr string) {
	p.address = adr
}
