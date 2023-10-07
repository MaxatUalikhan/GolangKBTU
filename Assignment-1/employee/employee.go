package employee

type Employee interface {
	GetPosition() string
	GetSalary() float64
	GetAddress() string
	SetPosition(string)
	SetSalary(float64)
	SetAddress(string)
}
